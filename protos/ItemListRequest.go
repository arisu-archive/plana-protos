package protos

type ItemListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
