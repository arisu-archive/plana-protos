package protos

type EventImageRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventId int64 `json:",omitempty,omitzero"`
}
