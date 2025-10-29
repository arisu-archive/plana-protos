package protos

type AccountCallNameRequest struct {
	RequestPacket
	CallName string `json:",omitempty,omitzero"`
}
