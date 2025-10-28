package protos

type EventListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
