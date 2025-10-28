package protos

type MiniGameCCGEnterStageRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	NodeId int64 `json:",omitempty,omitzero"`
}
