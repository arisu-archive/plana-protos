package protos

type CraftShiftingCompleteProcessResponse struct {
	ResponsePacket
	CraftInfoDB    ShiftingCraftInfoDB
	ParcelResultDB ParcelResultDB
}
