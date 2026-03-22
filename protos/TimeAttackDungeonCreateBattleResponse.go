package protos

type TimeAttackDungeonCreateBattleResponse struct {
	ResponsePacket
	RoomDB         TimeAttackDungeonRoomDB
	ParcelResultDB ParcelResultDB
}
