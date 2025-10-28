package protos

type CafeClearPresetResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
