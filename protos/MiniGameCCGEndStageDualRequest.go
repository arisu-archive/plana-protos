package protos

type MiniGameCCGEndStageDualRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	Summary MiniGameCCGSummary `json:",omitempty,omitzero"`
}
