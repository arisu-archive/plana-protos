package protos

type SkipHistoryListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SkipHistoryDB SkipHistoryDB `json:",omitempty,omitzero"`
}
