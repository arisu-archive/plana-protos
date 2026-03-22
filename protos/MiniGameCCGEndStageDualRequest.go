package protos

type MiniGameCCGEndStageDualRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	Summary        MiniGameCCGSummary
}
