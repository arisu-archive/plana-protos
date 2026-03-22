package protos

type TimeAttackDungeonSweepResponse struct {
	ResponsePacket
	Rewards        [][]ParcelInfo
	ParcelResultDB ParcelResultDB
	RoomDB         TimeAttackDungeonRoomDB
}
