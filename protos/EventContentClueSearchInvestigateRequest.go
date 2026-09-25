package protos

type EventContentClueSearchInvestigateRequest struct {
	RequestPacket
	EventContentId   int64 `json:",omitempty,omitzero"`
	InvestigateCount int32 `json:",omitempty,omitzero"`
}
