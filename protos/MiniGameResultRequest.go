package protos

type MiniGameResultRequest struct {
	RequestPacket
	EventContentId int64                 `json:",omitempty,omitzero"`
	UniqueId       int64                 `json:",omitempty,omitzero"`
	Summary        MinigameRhythmSummary `json:",omitempty,omitzero"`
}
