package protos

type EventContentClueSearchSubmitRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	ClueId         int64 `json:",omitempty,omitzero"`
}
