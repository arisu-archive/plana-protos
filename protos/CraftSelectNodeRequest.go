package protos

type CraftSelectNodeRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SlotId int64 `json:",omitempty,omitzero"`
	LeafNodeIndex int64 `json:",omitempty,omitzero"`
}
