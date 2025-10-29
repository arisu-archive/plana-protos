package protos

type CraftRewardRequest struct {
	RequestPacket
	SlotId int64 `json:",omitempty,omitzero"`
}
