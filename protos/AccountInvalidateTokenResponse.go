package protos

type AccountInvalidateTokenResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
