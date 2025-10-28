package protos

type ScenarioListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ScenarioHistoryDBs []ScenarioHistoryDB `json:",omitempty,omitzero"`
	ScenarioGroupHistoryDBs []ScenarioGroupHistoryDB `json:",omitempty,omitzero"`
	ScenarioCollectionDBs []ScenarioCollectionDB `json:",omitempty,omitzero"`
}
