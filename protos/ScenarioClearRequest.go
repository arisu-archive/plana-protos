package protos

type ScenarioClearRequest struct {
	RequestPacket
	ScenarioId    int64         `json:",omitempty,omitzero"`
	BattleSummary BattleSummary `json:",omitempty,omitzero"`
}
