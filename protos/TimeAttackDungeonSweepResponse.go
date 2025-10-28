package protos

type TimeAttackDungeonSweepResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Rewards [][]ParcelInfo `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	RoomDB TimeAttackDungeonRoomDB `json:",omitempty,omitzero"`
}
