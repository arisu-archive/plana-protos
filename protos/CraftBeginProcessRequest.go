package protos

type CraftBeginProcessRequest struct {
	RequestPacket
	SlotId int64 `json:",omitempty,omitzero"`
}
