package protos

type CraftSelectNodeRequest struct {
	RequestPacket
	SlotId int64 `json:",omitempty,omitzero"`
	LeafNodeIndex int64 `json:",omitempty,omitzero"`
}
