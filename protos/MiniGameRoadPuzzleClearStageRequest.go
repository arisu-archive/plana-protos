package protos

type MiniGameRoadPuzzleClearStageRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	UniqueId       int64 `json:",omitempty,omitzero"`
	Round          int64 `json:",omitempty,omitzero"`
	IsSkip         bool  `json:",omitempty,omitzero"`
}
