package protos

type ScenarioGroupHistoryUpdateResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ScenarioGroupHistoryDB ScenarioGroupHistoryDB `json:",omitempty,omitzero"`
}
