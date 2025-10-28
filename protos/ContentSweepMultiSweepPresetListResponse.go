package protos

type ContentSweepMultiSweepPresetListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MultiSweepPresetDBs []MultiSweepPresetDB `json:",omitempty,omitzero"`
}
