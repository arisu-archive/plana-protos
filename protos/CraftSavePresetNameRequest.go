package protos

type CraftSavePresetNameRequest struct {
	RequestPacket
	PresetIndex int32 `json:",omitempty,omitzero"`
	PresetName  string
}
