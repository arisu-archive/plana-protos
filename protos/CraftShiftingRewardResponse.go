package protos

type CraftShiftingRewardResponse struct {
	ResponsePacket
	ParcelResultDB   ParcelResultDB
	TargetCraftInfos []ShiftingCraftInfoDB
}
