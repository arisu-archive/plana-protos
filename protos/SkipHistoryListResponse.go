package protos

type SkipHistoryListResponse struct {
	ResponsePacket
	SkipHistoryDB SkipHistoryDB `json:",omitempty,omitzero"`
}
