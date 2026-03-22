package protos

type CraftShiftingBeginProcessResponse struct {
	ResponsePacket
	CraftInfoDB    ShiftingCraftInfoDB
	ParcelResultDB ParcelResultDB
}
