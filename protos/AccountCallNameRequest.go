package protos

type AccountCallNameRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CallName string `json:",omitempty,omitzero"`
}
