package protos

type CraftShiftingCompleteProcessRequest struct {
	RequestPacket
	SlotId int64 `json:",omitempty,omitzero"`
}
