package protos

type CafeListPresetResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CafePresetDBs []CafePresetDB `json:",omitempty,omitzero"`
}
