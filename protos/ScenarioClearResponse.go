package protos

type ScenarioClearResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ScenarioHistoryDB ScenarioHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ScenarioCollectionDBs []ScenarioCollectionDB `json:",omitempty,omitzero"`
}
