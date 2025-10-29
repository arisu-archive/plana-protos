package protos

type CraftShiftingRewardResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	TargetCraftInfos []ShiftingCraftInfoDB `json:",omitempty,omitzero"`
}
