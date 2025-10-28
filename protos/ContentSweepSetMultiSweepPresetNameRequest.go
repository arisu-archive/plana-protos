package protos

type ContentSweepSetMultiSweepPresetNameRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PresetId int64 `json:",omitempty,omitzero"`
	PresetName string `json:",omitempty,omitzero"`
}
