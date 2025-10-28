package protos

type ResetableContentGetRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
