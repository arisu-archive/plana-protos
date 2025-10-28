package protos

type SkipHistorySaveResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SkipHistoryDB SkipHistoryDB `json:",omitempty,omitzero"`
}
