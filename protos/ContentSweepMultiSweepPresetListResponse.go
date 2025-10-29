package protos

type ContentSweepMultiSweepPresetListResponse struct {
	ResponsePacket
	MultiSweepPresetDBs []MultiSweepPresetDB `json:",omitempty,omitzero"`
}
