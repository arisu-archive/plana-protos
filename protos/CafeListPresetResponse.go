package protos

type CafeListPresetResponse struct {
	ResponsePacket
	CafePresetDBs     []CafePresetDB `json:",omitempty,omitzero"`
	CafeCopyPresetDBs []CafePresetDB `json:",omitempty,omitzero"`
}
