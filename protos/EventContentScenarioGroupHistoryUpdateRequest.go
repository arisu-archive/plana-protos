package protos

type EventContentScenarioGroupHistoryUpdateRequest struct {
	RequestPacket
	ScenarioGroupUniqueId int64 `json:",omitempty,omitzero"`
	ScenarioType          int64 `json:",omitempty,omitzero"`
	EventContentId        int64 `json:",omitempty,omitzero"`
}
