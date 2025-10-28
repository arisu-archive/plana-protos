package protos

type AccountInvalidateTokenRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
