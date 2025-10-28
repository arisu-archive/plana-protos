package protos

type SkipHistoryListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
