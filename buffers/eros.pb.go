// Code generated by protoc-gen-go.
// source: buffers/eros.proto
// DO NOT EDIT!

/*
Package protobufs is a generated protocol buffer package.

It is generated from these files:
	buffers/eros.proto

It has these top-level messages:
	Handshake
	Division
	HandshakeResponse
	UserRegionStats
	UserStats
	MapPool
	Map
	SimulationResult
	MatchmakingQueue
	MatchmakingResult
	ChatRoomInfo
	ChatRoomIndex
	ChatMessage
	ChatRoomMessage
	ChatPrivateMessage
	ChatRoomUser
	ChatRoomRequest
	MatchmakingStats
	ServerStats
	Character
	OAuthRequest
	OAuthUrl
	MatchParticipant
	MatchResult
	BroadcastAlert
*/
package protobufs

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Region int32

const (
	Region_NA  Region = 1
	Region_EU  Region = 2
	Region_KR  Region = 3
	Region_CN  Region = 5
	Region_SEA Region = 6
)

var Region_name = map[int32]string{
	1: "NA",
	2: "EU",
	3: "KR",
	5: "CN",
	6: "SEA",
}
var Region_value = map[string]int32{
	"NA":  1,
	"EU":  2,
	"KR":  3,
	"CN":  5,
	"SEA": 6,
}

func (x Region) Enum() *Region {
	p := new(Region)
	*p = x
	return p
}
func (x Region) String() string {
	return proto.EnumName(Region_name, int32(x))
}
func (x *Region) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Region_value, data, "Region")
	if err != nil {
		return err
	}
	*x = Region(value)
	return nil
}

type HandshakeResponse_HandshakeStatus int32

const (
	HandshakeResponse_FAIL              HandshakeResponse_HandshakeStatus = 0
	HandshakeResponse_SUCCESS           HandshakeResponse_HandshakeStatus = 1
	HandshakeResponse_ALREADY_LOGGED_IN HandshakeResponse_HandshakeStatus = 2
)

var HandshakeResponse_HandshakeStatus_name = map[int32]string{
	0: "FAIL",
	1: "SUCCESS",
	2: "ALREADY_LOGGED_IN",
}
var HandshakeResponse_HandshakeStatus_value = map[string]int32{
	"FAIL":              0,
	"SUCCESS":           1,
	"ALREADY_LOGGED_IN": 2,
}

func (x HandshakeResponse_HandshakeStatus) Enum() *HandshakeResponse_HandshakeStatus {
	p := new(HandshakeResponse_HandshakeStatus)
	*p = x
	return p
}
func (x HandshakeResponse_HandshakeStatus) String() string {
	return proto.EnumName(HandshakeResponse_HandshakeStatus_name, int32(x))
}
func (x *HandshakeResponse_HandshakeStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(HandshakeResponse_HandshakeStatus_value, data, "HandshakeResponse_HandshakeStatus")
	if err != nil {
		return err
	}
	*x = HandshakeResponse_HandshakeStatus(value)
	return nil
}

type Handshake struct {
	Username         *string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	AuthKey          *string `protobuf:"bytes,2,opt,name=auth_key" json:"auth_key,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Handshake) Reset()         { *m = Handshake{} }
func (m *Handshake) String() string { return proto.CompactTextString(m) }
func (*Handshake) ProtoMessage()    {}

func (m *Handshake) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *Handshake) GetAuthKey() string {
	if m != nil && m.AuthKey != nil {
		return *m.AuthKey
	}
	return ""
}

type Division struct {
	Id                 *int64   `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Name               *string  `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	PromotionThreshold *float64 `protobuf:"fixed64,3,req,name=promotion_threshold" json:"promotion_threshold,omitempty"`
	DemotionThreshold  *float64 `protobuf:"fixed64,4,req,name=demotion_threshold" json:"demotion_threshold,omitempty"`
	IconUrl            *string  `protobuf:"bytes,5,req,name=icon_url" json:"icon_url,omitempty"`
	SmallIconUrl       *string  `protobuf:"bytes,6,req,name=small_icon_url" json:"small_icon_url,omitempty"`
	XXX_unrecognized   []byte   `json:"-"`
}

func (m *Division) Reset()         { *m = Division{} }
func (m *Division) String() string { return proto.CompactTextString(m) }
func (*Division) ProtoMessage()    {}

func (m *Division) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Division) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Division) GetPromotionThreshold() float64 {
	if m != nil && m.PromotionThreshold != nil {
		return *m.PromotionThreshold
	}
	return 0
}

func (m *Division) GetDemotionThreshold() float64 {
	if m != nil && m.DemotionThreshold != nil {
		return *m.DemotionThreshold
	}
	return 0
}

func (m *Division) GetIconUrl() string {
	if m != nil && m.IconUrl != nil {
		return *m.IconUrl
	}
	return ""
}

func (m *Division) GetSmallIconUrl() string {
	if m != nil && m.SmallIconUrl != nil {
		return *m.SmallIconUrl
	}
	return ""
}

type HandshakeResponse struct {
	Status           *HandshakeResponse_HandshakeStatus `protobuf:"varint,1,req,name=status,enum=protobufs.HandshakeResponse_HandshakeStatus" json:"status,omitempty"`
	User             *UserStats                         `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Id               *int64                             `protobuf:"varint,3,opt,name=id" json:"id,omitempty"`
	Character        []*Character                       `protobuf:"bytes,4,rep,name=character" json:"character,omitempty"`
	Division         []*Division                        `protobuf:"bytes,5,rep,name=division" json:"division,omitempty"`
	ActiveRegion     []Region                           `protobuf:"varint,6,rep,name=active_region,enum=protobufs.Region" json:"active_region,omitempty"`
	MapPool          *MapPool                           `protobuf:"bytes,7,opt,name=map_pool" json:"map_pool,omitempty"`
	MaxVetoes        *int64                             `protobuf:"varint,8,opt,name=max_vetoes" json:"max_vetoes,omitempty"`
	XXX_unrecognized []byte                             `json:"-"`
}

func (m *HandshakeResponse) Reset()         { *m = HandshakeResponse{} }
func (m *HandshakeResponse) String() string { return proto.CompactTextString(m) }
func (*HandshakeResponse) ProtoMessage()    {}

func (m *HandshakeResponse) GetStatus() HandshakeResponse_HandshakeStatus {
	if m != nil && m.Status != nil {
		return *m.Status
	}
	return HandshakeResponse_FAIL
}

func (m *HandshakeResponse) GetUser() *UserStats {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *HandshakeResponse) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *HandshakeResponse) GetCharacter() []*Character {
	if m != nil {
		return m.Character
	}
	return nil
}

func (m *HandshakeResponse) GetDivision() []*Division {
	if m != nil {
		return m.Division
	}
	return nil
}

func (m *HandshakeResponse) GetActiveRegion() []Region {
	if m != nil {
		return m.ActiveRegion
	}
	return nil
}

func (m *HandshakeResponse) GetMapPool() *MapPool {
	if m != nil {
		return m.MapPool
	}
	return nil
}

func (m *HandshakeResponse) GetMaxVetoes() int64 {
	if m != nil && m.MaxVetoes != nil {
		return *m.MaxVetoes
	}
	return 0
}

type UserRegionStats struct {
	Region              *Region  `protobuf:"varint,1,req,name=region,enum=protobufs.Region" json:"region,omitempty"`
	Points              *int64   `protobuf:"varint,2,req,name=points" json:"points,omitempty"`
	Wins                *int64   `protobuf:"varint,3,req,name=wins" json:"wins,omitempty"`
	Losses              *int64   `protobuf:"varint,4,req,name=losses" json:"losses,omitempty"`
	Forfeits            *int64   `protobuf:"varint,5,req,name=forfeits" json:"forfeits,omitempty"`
	Walkovers           *int64   `protobuf:"varint,6,req,name=walkovers" json:"walkovers,omitempty"`
	Mmr                 *float64 `protobuf:"fixed64,7,req,name=mmr" json:"mmr,omitempty"`
	PlacementsRemaining *int64   `protobuf:"varint,8,req,name=placements_remaining" json:"placements_remaining,omitempty"`
	Division            *int64   `protobuf:"varint,9,req,name=division" json:"division,omitempty"`
	DivisionRank        *int64   `protobuf:"varint,10,opt,name=division_rank" json:"division_rank,omitempty"`
	XXX_unrecognized    []byte   `json:"-"`
}

func (m *UserRegionStats) Reset()         { *m = UserRegionStats{} }
func (m *UserRegionStats) String() string { return proto.CompactTextString(m) }
func (*UserRegionStats) ProtoMessage()    {}

func (m *UserRegionStats) GetRegion() Region {
	if m != nil && m.Region != nil {
		return *m.Region
	}
	return Region_NA
}

func (m *UserRegionStats) GetPoints() int64 {
	if m != nil && m.Points != nil {
		return *m.Points
	}
	return 0
}

func (m *UserRegionStats) GetWins() int64 {
	if m != nil && m.Wins != nil {
		return *m.Wins
	}
	return 0
}

func (m *UserRegionStats) GetLosses() int64 {
	if m != nil && m.Losses != nil {
		return *m.Losses
	}
	return 0
}

func (m *UserRegionStats) GetForfeits() int64 {
	if m != nil && m.Forfeits != nil {
		return *m.Forfeits
	}
	return 0
}

func (m *UserRegionStats) GetWalkovers() int64 {
	if m != nil && m.Walkovers != nil {
		return *m.Walkovers
	}
	return 0
}

func (m *UserRegionStats) GetMmr() float64 {
	if m != nil && m.Mmr != nil {
		return *m.Mmr
	}
	return 0
}

func (m *UserRegionStats) GetPlacementsRemaining() int64 {
	if m != nil && m.PlacementsRemaining != nil {
		return *m.PlacementsRemaining
	}
	return 0
}

func (m *UserRegionStats) GetDivision() int64 {
	if m != nil && m.Division != nil {
		return *m.Division
	}
	return 0
}

func (m *UserRegionStats) GetDivisionRank() int64 {
	if m != nil && m.DivisionRank != nil {
		return *m.DivisionRank
	}
	return 0
}

type UserStats struct {
	Username            *string            `protobuf:"bytes,1,req,name=username" json:"username,omitempty"`
	SearchRadius        *int64             `protobuf:"varint,2,req,name=search_radius" json:"search_radius,omitempty"`
	Points              *int64             `protobuf:"varint,3,req,name=points" json:"points,omitempty"`
	Wins                *int64             `protobuf:"varint,4,req,name=wins" json:"wins,omitempty"`
	Losses              *int64             `protobuf:"varint,5,req,name=losses" json:"losses,omitempty"`
	Forfeits            *int64             `protobuf:"varint,6,req,name=forfeits" json:"forfeits,omitempty"`
	Walkovers           *int64             `protobuf:"varint,7,req,name=walkovers" json:"walkovers,omitempty"`
	Region              []*UserRegionStats `protobuf:"bytes,8,rep,name=region" json:"region,omitempty"`
	Vetoes              []*Map             `protobuf:"bytes,9,rep,name=vetoes" json:"vetoes,omitempty"`
	Id                  *int64             `protobuf:"varint,10,req,name=id" json:"id,omitempty"`
	Mmr                 *float64           `protobuf:"fixed64,11,req,name=mmr" json:"mmr,omitempty"`
	PlacementsRemaining *int64             `protobuf:"varint,12,req,name=placements_remaining" json:"placements_remaining,omitempty"`
	Division            *int64             `protobuf:"varint,13,req,name=division" json:"division,omitempty"`
	DivisionRank        *int64             `protobuf:"varint,14,req,name=division_rank" json:"division_rank,omitempty"`
	XXX_unrecognized    []byte             `json:"-"`
}

func (m *UserStats) Reset()         { *m = UserStats{} }
func (m *UserStats) String() string { return proto.CompactTextString(m) }
func (*UserStats) ProtoMessage()    {}

func (m *UserStats) GetUsername() string {
	if m != nil && m.Username != nil {
		return *m.Username
	}
	return ""
}

func (m *UserStats) GetSearchRadius() int64 {
	if m != nil && m.SearchRadius != nil {
		return *m.SearchRadius
	}
	return 0
}

func (m *UserStats) GetPoints() int64 {
	if m != nil && m.Points != nil {
		return *m.Points
	}
	return 0
}

func (m *UserStats) GetWins() int64 {
	if m != nil && m.Wins != nil {
		return *m.Wins
	}
	return 0
}

func (m *UserStats) GetLosses() int64 {
	if m != nil && m.Losses != nil {
		return *m.Losses
	}
	return 0
}

func (m *UserStats) GetForfeits() int64 {
	if m != nil && m.Forfeits != nil {
		return *m.Forfeits
	}
	return 0
}

func (m *UserStats) GetWalkovers() int64 {
	if m != nil && m.Walkovers != nil {
		return *m.Walkovers
	}
	return 0
}

func (m *UserStats) GetRegion() []*UserRegionStats {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *UserStats) GetVetoes() []*Map {
	if m != nil {
		return m.Vetoes
	}
	return nil
}

func (m *UserStats) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *UserStats) GetMmr() float64 {
	if m != nil && m.Mmr != nil {
		return *m.Mmr
	}
	return 0
}

func (m *UserStats) GetPlacementsRemaining() int64 {
	if m != nil && m.PlacementsRemaining != nil {
		return *m.PlacementsRemaining
	}
	return 0
}

func (m *UserStats) GetDivision() int64 {
	if m != nil && m.Division != nil {
		return *m.Division
	}
	return 0
}

func (m *UserStats) GetDivisionRank() int64 {
	if m != nil && m.DivisionRank != nil {
		return *m.DivisionRank
	}
	return 0
}

type MapPool struct {
	Map              []*Map `protobuf:"bytes,1,rep,name=map" json:"map,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *MapPool) Reset()         { *m = MapPool{} }
func (m *MapPool) String() string { return proto.CompactTextString(m) }
func (*MapPool) ProtoMessage()    {}

func (m *MapPool) GetMap() []*Map {
	if m != nil {
		return m.Map
	}
	return nil
}

type Map struct {
	Region           *Region `protobuf:"varint,1,req,name=region,enum=protobufs.Region" json:"region,omitempty"`
	BattleNetName    *string `protobuf:"bytes,2,req,name=battle_net_name" json:"battle_net_name,omitempty"`
	BattleNetId      *int32  `protobuf:"varint,3,req,name=battle_net_id" json:"battle_net_id,omitempty"`
	Description      *string `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	InfoUrl          *string `protobuf:"bytes,5,opt,name=info_url" json:"info_url,omitempty"`
	PreviewUrl       *string `protobuf:"bytes,6,opt,name=preview_url" json:"preview_url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Map) Reset()         { *m = Map{} }
func (m *Map) String() string { return proto.CompactTextString(m) }
func (*Map) ProtoMessage()    {}

func (m *Map) GetRegion() Region {
	if m != nil && m.Region != nil {
		return *m.Region
	}
	return Region_NA
}

func (m *Map) GetBattleNetName() string {
	if m != nil && m.BattleNetName != nil {
		return *m.BattleNetName
	}
	return ""
}

func (m *Map) GetBattleNetId() int32 {
	if m != nil && m.BattleNetId != nil {
		return *m.BattleNetId
	}
	return 0
}

func (m *Map) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *Map) GetInfoUrl() string {
	if m != nil && m.InfoUrl != nil {
		return *m.InfoUrl
	}
	return ""
}

func (m *Map) GetPreviewUrl() string {
	if m != nil && m.PreviewUrl != nil {
		return *m.PreviewUrl
	}
	return ""
}

type SimulationResult struct {
	Opponent         *UserStats `protobuf:"bytes,1,req,name=opponent" json:"opponent,omitempty"`
	Victory          *bool      `protobuf:"varint,2,req,name=victory" json:"victory,omitempty"`
	MatchQuality     *float64   `protobuf:"fixed64,3,req,name=match_quality" json:"match_quality,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *SimulationResult) Reset()         { *m = SimulationResult{} }
func (m *SimulationResult) String() string { return proto.CompactTextString(m) }
func (*SimulationResult) ProtoMessage()    {}

func (m *SimulationResult) GetOpponent() *UserStats {
	if m != nil {
		return m.Opponent
	}
	return nil
}

func (m *SimulationResult) GetVictory() bool {
	if m != nil && m.Victory != nil {
		return *m.Victory
	}
	return false
}

func (m *SimulationResult) GetMatchQuality() float64 {
	if m != nil && m.MatchQuality != nil {
		return *m.MatchQuality
	}
	return 0
}

type MatchmakingQueue struct {
	Region           []Region `protobuf:"varint,1,rep,name=region,enum=protobufs.Region" json:"region,omitempty"`
	Radius           *int64   `protobuf:"varint,2,req,name=radius" json:"radius,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *MatchmakingQueue) Reset()         { *m = MatchmakingQueue{} }
func (m *MatchmakingQueue) String() string { return proto.CompactTextString(m) }
func (*MatchmakingQueue) ProtoMessage()    {}

func (m *MatchmakingQueue) GetRegion() []Region {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *MatchmakingQueue) GetRadius() int64 {
	if m != nil && m.Radius != nil {
		return *m.Radius
	}
	return 0
}

type MatchmakingResult struct {
	Timespan         *int64     `protobuf:"varint,1,req,name=timespan" json:"timespan,omitempty"`
	Quality          *float64   `protobuf:"fixed64,2,req,name=quality" json:"quality,omitempty"`
	Opponent         *UserStats `protobuf:"bytes,3,req,name=opponent" json:"opponent,omitempty"`
	OpponentLatency  *int64     `protobuf:"varint,4,req,name=opponent_latency" json:"opponent_latency,omitempty"`
	Channel          *string    `protobuf:"bytes,5,req,name=channel" json:"channel,omitempty"`
	ChatRoom         *string    `protobuf:"bytes,6,req,name=chat_room" json:"chat_room,omitempty"`
	Map              *Map       `protobuf:"bytes,7,req,name=map" json:"map,omitempty"`
	LongUnlockTime   *int64     `protobuf:"varint,8,req,name=long_unlock_time" json:"long_unlock_time,omitempty"`
	LongResponseTime *int64     `protobuf:"varint,9,req,name=long_response_time" json:"long_response_time,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *MatchmakingResult) Reset()         { *m = MatchmakingResult{} }
func (m *MatchmakingResult) String() string { return proto.CompactTextString(m) }
func (*MatchmakingResult) ProtoMessage()    {}

func (m *MatchmakingResult) GetTimespan() int64 {
	if m != nil && m.Timespan != nil {
		return *m.Timespan
	}
	return 0
}

func (m *MatchmakingResult) GetQuality() float64 {
	if m != nil && m.Quality != nil {
		return *m.Quality
	}
	return 0
}

func (m *MatchmakingResult) GetOpponent() *UserStats {
	if m != nil {
		return m.Opponent
	}
	return nil
}

func (m *MatchmakingResult) GetOpponentLatency() int64 {
	if m != nil && m.OpponentLatency != nil {
		return *m.OpponentLatency
	}
	return 0
}

func (m *MatchmakingResult) GetChannel() string {
	if m != nil && m.Channel != nil {
		return *m.Channel
	}
	return ""
}

func (m *MatchmakingResult) GetChatRoom() string {
	if m != nil && m.ChatRoom != nil {
		return *m.ChatRoom
	}
	return ""
}

func (m *MatchmakingResult) GetMap() *Map {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *MatchmakingResult) GetLongUnlockTime() int64 {
	if m != nil && m.LongUnlockTime != nil {
		return *m.LongUnlockTime
	}
	return 0
}

func (m *MatchmakingResult) GetLongResponseTime() int64 {
	if m != nil && m.LongResponseTime != nil {
		return *m.LongResponseTime
	}
	return 0
}

type ChatRoomInfo struct {
	Key              *string      `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Name             *string      `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Passworded       *bool        `protobuf:"varint,3,req,name=passworded" json:"passworded,omitempty"`
	Joinable         *bool        `protobuf:"varint,4,req,name=joinable" json:"joinable,omitempty"`
	Fixed            *bool        `protobuf:"varint,5,req,name=fixed" json:"fixed,omitempty"`
	Users            *int64       `protobuf:"varint,6,req,name=users" json:"users,omitempty"`
	Participant      []*UserStats `protobuf:"bytes,7,rep,name=participant" json:"participant,omitempty"`
	Forced           *bool        `protobuf:"varint,8,req,name=forced" json:"forced,omitempty"`
	XXX_unrecognized []byte       `json:"-"`
}

func (m *ChatRoomInfo) Reset()         { *m = ChatRoomInfo{} }
func (m *ChatRoomInfo) String() string { return proto.CompactTextString(m) }
func (*ChatRoomInfo) ProtoMessage()    {}

func (m *ChatRoomInfo) GetKey() string {
	if m != nil && m.Key != nil {
		return *m.Key
	}
	return ""
}

func (m *ChatRoomInfo) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ChatRoomInfo) GetPassworded() bool {
	if m != nil && m.Passworded != nil {
		return *m.Passworded
	}
	return false
}

func (m *ChatRoomInfo) GetJoinable() bool {
	if m != nil && m.Joinable != nil {
		return *m.Joinable
	}
	return false
}

func (m *ChatRoomInfo) GetFixed() bool {
	if m != nil && m.Fixed != nil {
		return *m.Fixed
	}
	return false
}

func (m *ChatRoomInfo) GetUsers() int64 {
	if m != nil && m.Users != nil {
		return *m.Users
	}
	return 0
}

func (m *ChatRoomInfo) GetParticipant() []*UserStats {
	if m != nil {
		return m.Participant
	}
	return nil
}

func (m *ChatRoomInfo) GetForced() bool {
	if m != nil && m.Forced != nil {
		return *m.Forced
	}
	return false
}

type ChatRoomIndex struct {
	Room             []*ChatRoomInfo `protobuf:"bytes,1,rep,name=room" json:"room,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *ChatRoomIndex) Reset()         { *m = ChatRoomIndex{} }
func (m *ChatRoomIndex) String() string { return proto.CompactTextString(m) }
func (*ChatRoomIndex) ProtoMessage()    {}

func (m *ChatRoomIndex) GetRoom() []*ChatRoomInfo {
	if m != nil {
		return m.Room
	}
	return nil
}

type ChatMessage struct {
	Sender           *string `protobuf:"bytes,1,req,name=sender" json:"sender,omitempty"`
	Target           *string `protobuf:"bytes,2,req,name=target" json:"target,omitempty"`
	Message          *string `protobuf:"bytes,3,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChatMessage) Reset()         { *m = ChatMessage{} }
func (m *ChatMessage) String() string { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()    {}

func (m *ChatMessage) GetSender() string {
	if m != nil && m.Sender != nil {
		return *m.Sender
	}
	return ""
}

func (m *ChatMessage) GetTarget() string {
	if m != nil && m.Target != nil {
		return *m.Target
	}
	return ""
}

func (m *ChatMessage) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type ChatRoomMessage struct {
	Room             *ChatRoomInfo `protobuf:"bytes,1,req,name=room" json:"room,omitempty"`
	Sender           *UserStats    `protobuf:"bytes,2,req,name=sender" json:"sender,omitempty"`
	Message          *string       `protobuf:"bytes,3,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *ChatRoomMessage) Reset()         { *m = ChatRoomMessage{} }
func (m *ChatRoomMessage) String() string { return proto.CompactTextString(m) }
func (*ChatRoomMessage) ProtoMessage()    {}

func (m *ChatRoomMessage) GetRoom() *ChatRoomInfo {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *ChatRoomMessage) GetSender() *UserStats {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *ChatRoomMessage) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type ChatPrivateMessage struct {
	Sender           *UserStats `protobuf:"bytes,1,req,name=sender" json:"sender,omitempty"`
	Message          *string    `protobuf:"bytes,2,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *ChatPrivateMessage) Reset()         { *m = ChatPrivateMessage{} }
func (m *ChatPrivateMessage) String() string { return proto.CompactTextString(m) }
func (*ChatPrivateMessage) ProtoMessage()    {}

func (m *ChatPrivateMessage) GetSender() *UserStats {
	if m != nil {
		return m.Sender
	}
	return nil
}

func (m *ChatPrivateMessage) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type ChatRoomUser struct {
	Room             *ChatRoomInfo `protobuf:"bytes,1,req,name=room" json:"room,omitempty"`
	User             *UserStats    `protobuf:"bytes,2,req,name=user" json:"user,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *ChatRoomUser) Reset()         { *m = ChatRoomUser{} }
func (m *ChatRoomUser) String() string { return proto.CompactTextString(m) }
func (*ChatRoomUser) ProtoMessage()    {}

func (m *ChatRoomUser) GetRoom() *ChatRoomInfo {
	if m != nil {
		return m.Room
	}
	return nil
}

func (m *ChatRoomUser) GetUser() *UserStats {
	if m != nil {
		return m.User
	}
	return nil
}

type ChatRoomRequest struct {
	Room             *string `protobuf:"bytes,1,req,name=room" json:"room,omitempty"`
	Password         *string `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ChatRoomRequest) Reset()         { *m = ChatRoomRequest{} }
func (m *ChatRoomRequest) String() string { return proto.CompactTextString(m) }
func (*ChatRoomRequest) ProtoMessage()    {}

func (m *ChatRoomRequest) GetRoom() string {
	if m != nil && m.Room != nil {
		return *m.Room
	}
	return ""
}

func (m *ChatRoomRequest) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

type MatchmakingStats struct {
	Region           *Region `protobuf:"varint,1,req,name=region,enum=protobufs.Region" json:"region,omitempty"`
	SearchingUsers   *int64  `protobuf:"varint,2,req,name=searching_users" json:"searching_users,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *MatchmakingStats) Reset()         { *m = MatchmakingStats{} }
func (m *MatchmakingStats) String() string { return proto.CompactTextString(m) }
func (*MatchmakingStats) ProtoMessage()    {}

func (m *MatchmakingStats) GetRegion() Region {
	if m != nil && m.Region != nil {
		return *m.Region
	}
	return Region_NA
}

func (m *MatchmakingStats) GetSearchingUsers() int64 {
	if m != nil && m.SearchingUsers != nil {
		return *m.SearchingUsers
	}
	return 0
}

type ServerStats struct {
	ActiveUsers      *int64              `protobuf:"varint,1,req,name=active_users" json:"active_users,omitempty"`
	SearchingUsers   *int64              `protobuf:"varint,2,req,name=searching_users" json:"searching_users,omitempty"`
	Region           []*MatchmakingStats `protobuf:"bytes,3,rep,name=region" json:"region,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *ServerStats) Reset()         { *m = ServerStats{} }
func (m *ServerStats) String() string { return proto.CompactTextString(m) }
func (*ServerStats) ProtoMessage()    {}

func (m *ServerStats) GetActiveUsers() int64 {
	if m != nil && m.ActiveUsers != nil {
		return *m.ActiveUsers
	}
	return 0
}

func (m *ServerStats) GetSearchingUsers() int64 {
	if m != nil && m.SearchingUsers != nil {
		return *m.SearchingUsers
	}
	return 0
}

func (m *ServerStats) GetRegion() []*MatchmakingStats {
	if m != nil {
		return m.Region
	}
	return nil
}

type Character struct {
	Region               *Region `protobuf:"varint,1,req,name=region,enum=protobufs.Region" json:"region,omitempty"`
	Subregion            *int32  `protobuf:"varint,2,req,name=subregion" json:"subregion,omitempty"`
	ProfileId            *int32  `protobuf:"varint,3,req,name=profile_id" json:"profile_id,omitempty"`
	CharacterName        *string `protobuf:"bytes,4,req,name=character_name" json:"character_name,omitempty"`
	CharacterCode        *int32  `protobuf:"varint,5,opt,name=character_code" json:"character_code,omitempty"`
	ProfileLink          *string `protobuf:"bytes,6,opt,name=profile_link" json:"profile_link,omitempty"`
	IngameProfileLink    *string `protobuf:"bytes,7,opt,name=ingame_profile_link" json:"ingame_profile_link,omitempty"`
	Verified             *bool   `protobuf:"varint,8,opt,name=verified" json:"verified,omitempty"`
	VerificationPortrait *int32  `protobuf:"varint,9,opt,name=verification_portrait" json:"verification_portrait,omitempty"`
	XXX_unrecognized     []byte  `json:"-"`
}

func (m *Character) Reset()         { *m = Character{} }
func (m *Character) String() string { return proto.CompactTextString(m) }
func (*Character) ProtoMessage()    {}

func (m *Character) GetRegion() Region {
	if m != nil && m.Region != nil {
		return *m.Region
	}
	return Region_NA
}

func (m *Character) GetSubregion() int32 {
	if m != nil && m.Subregion != nil {
		return *m.Subregion
	}
	return 0
}

func (m *Character) GetProfileId() int32 {
	if m != nil && m.ProfileId != nil {
		return *m.ProfileId
	}
	return 0
}

func (m *Character) GetCharacterName() string {
	if m != nil && m.CharacterName != nil {
		return *m.CharacterName
	}
	return ""
}

func (m *Character) GetCharacterCode() int32 {
	if m != nil && m.CharacterCode != nil {
		return *m.CharacterCode
	}
	return 0
}

func (m *Character) GetProfileLink() string {
	if m != nil && m.ProfileLink != nil {
		return *m.ProfileLink
	}
	return ""
}

func (m *Character) GetIngameProfileLink() string {
	if m != nil && m.IngameProfileLink != nil {
		return *m.IngameProfileLink
	}
	return ""
}

func (m *Character) GetVerified() bool {
	if m != nil && m.Verified != nil {
		return *m.Verified
	}
	return false
}

func (m *Character) GetVerificationPortrait() int32 {
	if m != nil && m.VerificationPortrait != nil {
		return *m.VerificationPortrait
	}
	return 0
}

type OAuthRequest struct {
	Region           *Region `protobuf:"varint,1,req,name=region,enum=protobufs.Region" json:"region,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OAuthRequest) Reset()         { *m = OAuthRequest{} }
func (m *OAuthRequest) String() string { return proto.CompactTextString(m) }
func (*OAuthRequest) ProtoMessage()    {}

func (m *OAuthRequest) GetRegion() Region {
	if m != nil && m.Region != nil {
		return *m.Region
	}
	return Region_NA
}

type OAuthUrl struct {
	Url              *string `protobuf:"bytes,1,req,name=url" json:"url,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *OAuthUrl) Reset()         { *m = OAuthUrl{} }
func (m *OAuthUrl) String() string { return proto.CompactTextString(m) }
func (*OAuthUrl) ProtoMessage()    {}

func (m *OAuthUrl) GetUrl() string {
	if m != nil && m.Url != nil {
		return *m.Url
	}
	return ""
}

type MatchParticipant struct {
	User             *UserStats `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Character        *Character `protobuf:"bytes,2,opt,name=character" json:"character,omitempty"`
	PointsBefore     *int64     `protobuf:"varint,3,req,name=points_before" json:"points_before,omitempty"`
	PointsAfter      *int64     `protobuf:"varint,4,req,name=points_after" json:"points_after,omitempty"`
	PointsDifference *int64     `protobuf:"varint,5,req,name=points_difference" json:"points_difference,omitempty"`
	Victory          *bool      `protobuf:"varint,6,req,name=victory" json:"victory,omitempty"`
	Race             *string    `protobuf:"bytes,7,req,name=race" json:"race,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *MatchParticipant) Reset()         { *m = MatchParticipant{} }
func (m *MatchParticipant) String() string { return proto.CompactTextString(m) }
func (*MatchParticipant) ProtoMessage()    {}

func (m *MatchParticipant) GetUser() *UserStats {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *MatchParticipant) GetCharacter() *Character {
	if m != nil {
		return m.Character
	}
	return nil
}

func (m *MatchParticipant) GetPointsBefore() int64 {
	if m != nil && m.PointsBefore != nil {
		return *m.PointsBefore
	}
	return 0
}

func (m *MatchParticipant) GetPointsAfter() int64 {
	if m != nil && m.PointsAfter != nil {
		return *m.PointsAfter
	}
	return 0
}

func (m *MatchParticipant) GetPointsDifference() int64 {
	if m != nil && m.PointsDifference != nil {
		return *m.PointsDifference
	}
	return 0
}

func (m *MatchParticipant) GetVictory() bool {
	if m != nil && m.Victory != nil {
		return *m.Victory
	}
	return false
}

func (m *MatchParticipant) GetRace() string {
	if m != nil && m.Race != nil {
		return *m.Race
	}
	return ""
}

type MatchResult struct {
	Region           *Region             `protobuf:"varint,1,req,name=region,enum=protobufs.Region" json:"region,omitempty"`
	Map              *Map                `protobuf:"bytes,2,req,name=map" json:"map,omitempty"`
	Participant      []*MatchParticipant `protobuf:"bytes,3,rep,name=participant" json:"participant,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *MatchResult) Reset()         { *m = MatchResult{} }
func (m *MatchResult) String() string { return proto.CompactTextString(m) }
func (*MatchResult) ProtoMessage()    {}

func (m *MatchResult) GetRegion() Region {
	if m != nil && m.Region != nil {
		return *m.Region
	}
	return Region_NA
}

func (m *MatchResult) GetMap() *Map {
	if m != nil {
		return m.Map
	}
	return nil
}

func (m *MatchResult) GetParticipant() []*MatchParticipant {
	if m != nil {
		return m.Participant
	}
	return nil
}

type BroadcastAlert struct {
	Message          *string `protobuf:"bytes,2,req,name=message" json:"message,omitempty"`
	Predefined       *int32  `protobuf:"varint,1,req,name=predefined" json:"predefined,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BroadcastAlert) Reset()         { *m = BroadcastAlert{} }
func (m *BroadcastAlert) String() string { return proto.CompactTextString(m) }
func (*BroadcastAlert) ProtoMessage()    {}

func (m *BroadcastAlert) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *BroadcastAlert) GetPredefined() int32 {
	if m != nil && m.Predefined != nil {
		return *m.Predefined
	}
	return 0
}

func init() {
	proto.RegisterEnum("protobufs.Region", Region_name, Region_value)
	proto.RegisterEnum("protobufs.HandshakeResponse_HandshakeStatus", HandshakeResponse_HandshakeStatus_name, HandshakeResponse_HandshakeStatus_value)
}
