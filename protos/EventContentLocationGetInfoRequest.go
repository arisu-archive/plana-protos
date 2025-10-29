package protos

type EventContentLocationGetInfoRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
