package protos

type TimeAttackDungeonEndBattleResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RoomDB TimeAttackDungeonRoomDB `json:",omitempty,omitzero"`
	TotalPoint int64 `json:",omitempty,omitzero"`
	DefaultPoint int64 `json:",omitempty,omitzero"`
	TimePoint int64 `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
