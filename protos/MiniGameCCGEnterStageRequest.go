package protos

type MiniGameCCGEnterStageRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	NodeId int64 `json:",omitempty,omitzero"`
}
