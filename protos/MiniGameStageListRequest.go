package protos

type MiniGameStageListRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
