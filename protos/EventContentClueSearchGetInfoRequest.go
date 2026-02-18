package protos

type EventContentClueSearchGetInfoRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
