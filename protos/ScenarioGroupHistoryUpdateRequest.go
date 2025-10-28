package protos

type ScenarioGroupHistoryUpdateRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ScenarioGroupUniqueId int64 `json:",omitempty,omitzero"`
	ScenarioType int64 `json:",omitempty,omitzero"`
}
