package protos

type ContentSweepSetMultiSweepPresetNameResponse struct {
	ResponsePacket
	MultiSweepPresetDBs []MultiSweepPresetDB `json:",omitempty,omitzero"`
}
