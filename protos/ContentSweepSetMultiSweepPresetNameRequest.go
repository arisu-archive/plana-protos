package protos

type ContentSweepSetMultiSweepPresetNameRequest struct {
	RequestPacket
	PresetId int64 `json:",omitempty,omitzero"`
	PresetName string `json:",omitempty,omitzero"`
}
