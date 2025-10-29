package protos

type EventContentMainGroundStageResultRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	PassCheckCharacter bool `json:",omitempty,omitzero"`
	Summary BattleSummary `json:",omitempty,omitzero"`
}
