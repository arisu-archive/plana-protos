package protos

type ContentSweepSetMultiSweepPresetRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PresetId int64 `json:",omitempty,omitzero"`
	PresetName string `json:",omitempty,omitzero"`
	StageIds []int64 `json:",omitempty,omitzero"`
	ParcelIds []ParcelKeyPair `json:",omitempty,omitzero"`
}
