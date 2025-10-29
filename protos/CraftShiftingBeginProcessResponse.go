package protos

type CraftShiftingBeginProcessResponse struct {
	ResponsePacket
	CraftInfoDB ShiftingCraftInfoDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
