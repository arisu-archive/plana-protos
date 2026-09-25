package protos

type EventContentClueSearchRevealRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	ClueId         int64 `json:",omitempty,omitzero"`
}
