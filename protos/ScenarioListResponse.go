package protos

type ScenarioListResponse struct {
	ResponsePacket
	ScenarioHistoryDBs []ScenarioHistoryDB `json:",omitempty,omitzero"`
	ScenarioGroupHistoryDBs []ScenarioGroupHistoryDB `json:",omitempty,omitzero"`
	ScenarioCollectionDBs []ScenarioCollectionDB `json:",omitempty,omitzero"`
}
