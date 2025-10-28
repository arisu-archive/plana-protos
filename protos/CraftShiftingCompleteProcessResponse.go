package protos

type CraftShiftingCompleteProcessResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CraftInfoDB ShiftingCraftInfoDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
