package protos

type MiniGameResultRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	Summary MinigameRhythmSummary `json:",omitempty,omitzero"`
}
