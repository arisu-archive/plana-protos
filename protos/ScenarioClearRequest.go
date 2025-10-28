package protos

type ScenarioClearRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ScenarioId int64 `json:",omitempty,omitzero"`
	BattleSummary BattleSummary `json:",omitempty,omitzero"`
}
