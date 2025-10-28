package protos

type ContentSweepSetMultiSweepPresetNameResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MultiSweepPresetDBs []MultiSweepPresetDB `json:",omitempty,omitzero"`
}
