package protos

type ContentSweepSetMultiSweepPresetResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MultiSweepPresetDBs []MultiSweepPresetDB `json:",omitempty,omitzero"`
}
