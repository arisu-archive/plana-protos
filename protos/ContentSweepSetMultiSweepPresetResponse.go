package protos

type ContentSweepSetMultiSweepPresetResponse struct {
	ResponsePacket
	MultiSweepPresetDBs []MultiSweepPresetDB `json:",omitempty,omitzero"`
}
