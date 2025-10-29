package protos

type CraftCompleteProcessRequest struct {
	RequestPacket
	SlotId int64 `json:",omitempty,omitzero"`
}
