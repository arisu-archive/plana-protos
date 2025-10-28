package protos

type EventListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventInfoDBs []EventInfoDB `json:",omitempty,omitzero"`
}
