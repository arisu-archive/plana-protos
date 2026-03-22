package protos

type TimeAttackDungeonEndBattleResponse struct {
	ResponsePacket
	RoomDB         TimeAttackDungeonRoomDB
	TotalPoint     int64 `json:",omitempty,omitzero"`
	DefaultPoint   int64 `json:",omitempty,omitzero"`
	TimePoint      int64 `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB
}
