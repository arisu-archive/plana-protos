package protos

type CafeClearPresetRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SlotId int32 `json:",omitempty,omitzero"`
}
