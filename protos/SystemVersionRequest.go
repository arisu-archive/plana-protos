package protos

type SystemVersionRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
