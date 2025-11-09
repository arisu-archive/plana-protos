package protos

type CraftShiftingCompleteProcessResponse struct {
	ResponsePacket
	CraftInfoDB    ShiftingCraftInfoDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB      `json:",omitempty,omitzero"`
}
