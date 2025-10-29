package protos

type CraftShiftingCompleteProcessAllResponse struct {
	ResponsePacket
	CraftInfoDBs []ShiftingCraftInfoDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
