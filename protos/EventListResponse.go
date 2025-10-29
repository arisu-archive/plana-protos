package protos

type EventListResponse struct {
	ResponsePacket
	EventInfoDBs []EventInfoDB `json:",omitempty,omitzero"`
}
