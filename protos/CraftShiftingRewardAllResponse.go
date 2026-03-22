package protos

type CraftShiftingRewardAllResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	CraftInfoDBs   []ShiftingCraftInfoDB
}
