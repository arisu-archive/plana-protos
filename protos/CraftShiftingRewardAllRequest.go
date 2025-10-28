package protos

type CraftShiftingRewardAllRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
