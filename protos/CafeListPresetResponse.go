package protos

type CafeListPresetResponse struct {
	ResponsePacket
	CafePresetDBs     []CafePresetDB
	CafeCopyPresetDBs []CafePresetDB
}
