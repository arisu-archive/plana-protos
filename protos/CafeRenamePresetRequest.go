package protos

type CafeRenamePresetRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SlotId int32 `json:",omitempty,omitzero"`
	PresetName string `json:",omitempty,omitzero"`
}
