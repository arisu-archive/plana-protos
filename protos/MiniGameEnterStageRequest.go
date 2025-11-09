package protos

type MiniGameEnterStageRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	UniqueId       int64 `json:",omitempty,omitzero"`
}
