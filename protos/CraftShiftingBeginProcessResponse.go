package protos

type CraftShiftingBeginProcessResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CraftInfoDB ShiftingCraftInfoDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
