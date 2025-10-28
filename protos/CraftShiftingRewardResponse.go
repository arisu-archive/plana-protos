package protos

type CraftShiftingRewardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	TargetCraftInfos []ShiftingCraftInfoDB `json:",omitempty,omitzero"`
}
