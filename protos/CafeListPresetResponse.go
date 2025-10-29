package protos

type CafeListPresetResponse struct {
	ResponsePacket
	CafePresetDBs []CafePresetDB `json:",omitempty,omitzero"`
}
