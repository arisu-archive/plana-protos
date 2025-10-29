package protos

type SkipHistorySaveRequest struct {
	RequestPacket
	SkipHistoryDB SkipHistoryDB `json:",omitempty,omitzero"`
}
