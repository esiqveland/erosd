package main

import (
	"github.com/golang/protobuf/proto"
	"errors"
	"fmt"
	"github.com/Starbow/erosd/buffers"
	"log"
	"os"
	"path"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

const (
	ChannelMsg  = "CHM"
	ChannelJoin = "CHJ"
)

var (
	chatRooms                map[string]*ChatRoom
	joinableChatRooms        map[string]*ChatRoom
	fixedChatRooms           []string
	autoJoinChatRooms        []string
	maxChatRooms             int64 = 5
	chatIdBase               int64 = 1
	chatDelay                      = 250 * time.Millisecond
	chatMaxThrottleTime            = time.Duration(5 * time.Minute)
	chatMaxMessageLength     int64 = 256
	ErrChatRoomAlreadyExists error = errors.New("The chat room name specified already exists.")
	ErrChatRoomReserved      error = errors.New("The chat room name is reserved.")
	ErrChatRoomNameTooShort  error = errors.New("The chat room name is too short")
	_                              = log.Ldate
	maxMessageCache          int   = 2
)

func initChat() {
	chatRooms = make(map[string]*ChatRoom)
	joinableChatRooms = make(map[string]*ChatRoom)

	for x := range fixedChatRooms {
		room, err := NewChatRoom(fixedChatRooms[x], "", true, true)
		if err != nil {
			log.Println("Error creating channel", room, err)
		}
	}
}

type ChatRoom struct {
	id      int64
	members map[int64]*ClientConnection

	key      string // key identifying this chatroom
	name     string // Friendly name
	password string // Password
	joinable bool   // False if this is a server forced joining room (matchmaking)
	fixed    bool   // True if this is a server created room that never expires

	join    chan *ClientConnection
	leave   chan *ClientConnection
	message chan *protobufs.ChatRoomMessage
	abort   chan bool

	logger  *log.Logger
	logFile *os.File

	messageCache []*protobufs.ChatRoomMessage

	sync.RWMutex
}

func GetChatRoom(name string, password string, joinable, fixed bool) (room *ChatRoom) {
	name = cleanChatRoomName(name)
	room, ok := chatRooms[name]
	var err error
	if !ok {
		room, err = NewChatRoom(name, password, joinable, fixed)
		if err != nil {
			log.Println("Error creating chat", err, name)
		}
	}

	return room
}

// run is the main loop that runs and handles events for a ChatRoom.
func (cr *ChatRoom) run() {
	defer func() {
		// Handle closing the room
		delete(chatRooms, cr.key)
		delete(joinableChatRooms, cr.key)

		for x := range cr.members {
			delete(cr.members[x].chatRooms, cr.key)

		}
	}()
	timer := time.NewTicker(time.Second * 30)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			// Nobody has joined after a set time. Close the room if we can.
			if cr.CanShutdown() {
				cr.Close()
				return
			}

		case client := <-cr.join:
			// TODO: We might be able to do this in a go routine?
			cr.ClientJoin(client)

		case client := <-cr.leave:
			// TODO: We might be able to do this in a go routine?
			cr.ClientLeave(client)
			if cr.CanShutdown() {
				cr.Close()
				return
			}

		case msg := <-cr.message:
			cr.logger.Println("msg:", msg.GetSender().GetUsername(), ":", msg.GetMessage())
			cr.Broadcast(ChannelMsg, msg)
		case <-cr.abort:
			cr.logger.Println("Closing aborted room")
			cr.Close()
			return
		}
	}
}

func (cr *ChatRoom) ClientLeave(client *ClientConnection) {
	cr.Lock()
	defer cr.Unlock()

	_, exists := cr.members[client.id]
	if !exists {
		return
	}

	delete(cr.members, client.id)
	delete(client.chatRooms, cr.key)

	leave := cr.ChatRoomUserMessage(client, false)
	cr.Broadcast("CHL", &leave)
	cr.logger.Println("left:", client.id, client.client.Username)
}

// Broadcast a command and/or message to a ChatRoom minus exclude list
func (cr *ChatRoom) Broadcast(command string, message proto.Message, exclude ...*ClientConnection) error {
	data, err := Marshal(message)
	if err != nil {
		return err
	}

	cr.RLock()
	defer cr.RUnlock()

parent:
	for x := range cr.members {
		for y := range exclude {
			if cr.members[x] == exclude[y] {
				continue parent
			}
		}
		go cr.members[x].SendServerMessage(command, data)
	}

	return nil
}

func NewChatRoom(name, password string, joinable, fixed bool) (cr *ChatRoom, err error) {
	key := cleanChatRoomName(name)
	if len(key) < 3 {
		err = ErrChatRoomNameTooShort
		return
	}
	if joinable && key[:2] == "mm" {
		err = ErrChatRoomReserved
		return
	}

	_, ok := chatRooms[key]
	if ok {
		err = ErrChatRoomAlreadyExists
		return
	}

	cr = &ChatRoom{
		id:       atomic.AddInt64(&chatIdBase, 1),
		members:  make(map[int64]*ClientConnection),
		key:      key,
		name:     strings.TrimSpace(name),
		join:     make(chan *ClientConnection),
		leave:    make(chan *ClientConnection),
		message:  make(chan *protobufs.ChatRoomMessage, 256),
		abort:    make(chan bool, 1),
		joinable: joinable,
		fixed:    fixed,
		password: strings.TrimSpace(password),
	}

	var logfile string = path.Join(logPath, fmt.Sprintf("%d-chat-%d.log", os.Getpid(), cr.id))
	file, err := os.Create(logfile)
	if err != nil {
		log.Println("Failed to create log file", logfile, "for new chat", name)
		cr.logger = log.New(os.Stdout, fmt.Sprintf("chat-%d:", cr.id), log.Ldate|log.Ltime|log.Lshortfile)
	} else {
		log.Println("Logging new chat", cr.id, ":", name)
		cr.logger = log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
		cr.logFile = file
	}

	go cr.run()

	chatRooms[key] = cr
	if joinable {
		joinableChatRooms[key] = cr
	}
	return
}

func (ch *ChatRoom) ChatRoomInfoMessage(includeUserStats bool) *protobufs.ChatRoomInfo {
	ch.RLock()
	defer ch.RUnlock()
	var (
		chat       protobufs.ChatRoomInfo
		users      int64 = int64(len(ch.members))
		passworded bool  = ch.password != ""
		forced     bool  = false
	)

	for x := range autoJoinChatRooms {
		if autoJoinChatRooms[x] == ch.name {
			forced = true
			break
		}
	}

	chat.Key = &ch.key
	chat.Name = &ch.name

	chat.Passworded = &passworded
	chat.Fixed = &ch.fixed
	chat.Joinable = &ch.joinable
	chat.Users = &users
	chat.Forced = &forced

	if includeUserStats {
		chat.Participant = make([]*protobufs.UserStats, 0, users)
		for x := range ch.members {
			chat.Participant = append(chat.Participant, ch.members[x].client.UserStatsMessage())
		}
	}

	return &chat

}

func (ch *ChatRoom) ChatRoomUserMessage(client *ClientConnection, detailed bool) protobufs.ChatRoomUser {

	var (
		chat *protobufs.ChatRoomInfo = ch.ChatRoomInfoMessage(detailed)
		user *protobufs.UserStats    = client.client.UserStatsMessage()
		join protobufs.ChatRoomUser
	)
	join.Room = chat
	join.User = user

	return join
}

func (ch *ChatRoom) ChatRoomMessageMessage(client *ClientConnection, message *protobufs.ChatMessage) protobufs.ChatRoomMessage {

	var (
		chat *protobufs.ChatRoomInfo = ch.ChatRoomInfoMessage(false)
		user *protobufs.UserStats    = client.client.UserStatsMessage()
		msg  protobufs.ChatRoomMessage
		now  int64 = time.Now().Unix()
	)
	msg.Room = chat
	msg.Sender = user
	messageString := message.GetMessage()
	msg.Message = &messageString
	msg.Timestamp = &now

	// Save message into cache
	// This method will thrash the entire cache every time when full if we only delete message one by one...
	if len(ch.messageCache) >= maxMessageCache {
		ch.messageCache[0] = nil
		ch.messageCache = ch.messageCache[1:]
	}
	ch.messageCache = append(ch.messageCache, &msg)

	return msg

}

// canShutdown() checks if room can be shutdown
// We acquire the lock here, so make sure you don't already hold a writelock
func (cr *ChatRoom) CanShutdown() bool {
	cr.RLock()
	defer cr.RUnlock()

	if len(cr.members) != 0 || cr.fixed {
		return false
	}
	return true
}

// Close does cleanup so we can close the room
func (cr *ChatRoom) Close() {
	cr.logger.Printf("Closing room: %v", cr.key)

	if cr.logFile != nil {
		cr.logFile.Close()
	}
}

func (cr *ChatRoom) ClientJoin(client *ClientConnection) {
	if !cr.addMemberIfNew(client) {
		return
	}
	join := cr.ChatRoomUserMessage(client, false)
	cr.Broadcast(ChannelJoin, &join, client)
	detailedJoin := cr.ChatRoomUserMessage(client, true)
	data, _ := Marshal(&detailedJoin)
	client.SendServerMessage(ChannelJoin, data)
	client.chatRooms[cr.key] = cr
	cr.logger.Println("join:", client.id, client.client.Username)
}

func (cr *ChatRoom) addMemberIfNew(client *ClientConnection) bool {
	cr.Lock()
	defer cr.Unlock()

	_, exists := cr.members[client.id]
	if exists {
		return false
	}

	cr.members[client.id] = client
	return true
}

func cleanChatRoomName(name string) string {
	return strings.TrimSpace(strings.ToLower(name))
}
