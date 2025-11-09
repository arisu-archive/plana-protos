package protos

type ScenarioGroupHistoryUpdateRequest struct {
	RequestPacket
	ScenarioGroupUniqueId int64 `json:",omitempty,omitzero"`
	ScenarioType          int64 `json:",omitempty,omitzero"`
}
