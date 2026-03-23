package protos

type TimeAttackDungeonSweepResponse struct {
	ResponsePacket
	Rewards        [][]*ParcelInfo
	ParcelResultDB *ParcelResultDB          `json:",omitempty,omitzero"`
	RoomDB         *TimeAttackDungeonRoomDB `json:",omitempty,omitzero"`
}
