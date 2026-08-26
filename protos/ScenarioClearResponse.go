package protos

type ScenarioClearResponse struct {
	ResponsePacket
	ScenarioHistoryId     int64           `json:",omitempty,omitzero"`
	ParcelResultDB        *ParcelResultDB `json:",omitempty,omitzero"`
	ScenarioCollectionDBs []*ScenarioCollectionDB
}
