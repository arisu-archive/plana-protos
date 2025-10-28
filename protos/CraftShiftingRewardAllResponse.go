package protos

type CraftShiftingRewardAllResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	CraftInfoDBs []ShiftingCraftInfoDB `json:",omitempty,omitzero"`
}
