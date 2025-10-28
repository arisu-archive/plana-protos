package protos

type EventContentScenarioGroupHistoryUpdateRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ScenarioGroupUniqueId int64 `json:",omitempty,omitzero"`
	ScenarioType int64 `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
}
