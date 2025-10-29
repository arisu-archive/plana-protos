package protos

type TimeAttackDungeonSweepResponse struct {
	ResponsePacket
	Rewards [][]ParcelInfo `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	RoomDB TimeAttackDungeonRoomDB `json:",omitempty,omitzero"`
}
