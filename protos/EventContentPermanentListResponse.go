package protos

type EventContentPermanentListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PermanentDBs []EventContentPermanentDB `json:",omitempty,omitzero"`
}
