package protos

type CraftShiftingRewardAllResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	CraftInfoDBs []ShiftingCraftInfoDB `json:",omitempty,omitzero"`
}
