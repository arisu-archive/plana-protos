package protos

type AccountResetResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
