package protos

type EventImageRequest struct {
	RequestPacket
	EventId int64 `json:",omitempty,omitzero"`
}
