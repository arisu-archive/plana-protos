package protos

type CraftShiftingRewardRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SlotId int64 `json:",omitempty,omitzero"`
}
