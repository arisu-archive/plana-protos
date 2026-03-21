package protos

type EventContentClueSearchRoundCompleteRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
