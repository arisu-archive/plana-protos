package protos

type EventContentSubStageResultRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	PassCheckCharacter bool `json:",omitempty,omitzero"`
	Summary BattleSummary `json:",omitempty,omitzero"`
}
