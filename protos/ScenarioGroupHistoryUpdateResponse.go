package protos

type ScenarioGroupHistoryUpdateResponse struct {
	ResponsePacket
	ScenarioGroupHistoryDB ScenarioGroupHistoryDB `json:",omitempty,omitzero"`
}
