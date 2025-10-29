package protos

type CafeClearPresetRequest struct {
	RequestPacket
	SlotId int32 `json:",omitempty,omitzero"`
}
