package protos

type SkipHistorySaveResponse struct {
	ResponsePacket
	SkipHistoryDB SkipHistoryDB `json:",omitempty,omitzero"`
}
