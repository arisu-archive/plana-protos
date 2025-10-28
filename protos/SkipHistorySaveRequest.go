package protos

type SkipHistorySaveRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SkipHistoryDB SkipHistoryDB `json:",omitempty,omitzero"`
}
