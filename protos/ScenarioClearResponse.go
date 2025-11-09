package protos

type ScenarioClearResponse struct {
	ResponsePacket
	ScenarioHistoryDB     ScenarioHistoryDB      `json:",omitempty,omitzero"`
	ParcelResultDB        ParcelResultDB         `json:",omitempty,omitzero"`
	ScenarioCollectionDBs []ScenarioCollectionDB `json:",omitempty,omitzero"`
}
