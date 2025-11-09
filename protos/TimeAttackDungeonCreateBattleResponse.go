package protos

type TimeAttackDungeonCreateBattleResponse struct {
	ResponsePacket
	RoomDB         TimeAttackDungeonRoomDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB          `json:",omitempty,omitzero"`
}
