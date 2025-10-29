package protos

type CraftShiftingRewardRequest struct {
	RequestPacket
	SlotId int64 `json:",omitempty,omitzero"`
}
