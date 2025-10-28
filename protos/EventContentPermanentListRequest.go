package protos

type EventContentPermanentListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
