package protos

type CafeRenamePresetRequest struct {
	RequestPacket
	SlotId int32 `json:",omitempty,omitzero"`
	PresetName string `json:",omitempty,omitzero"`
}
