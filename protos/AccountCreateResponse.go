package protos

type AccountCreateResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
