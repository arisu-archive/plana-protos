package protos

type MiniGameCCGEndStageEventRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
