package protos

type ScenarioClearResponse struct {
	ResponsePacket
	ScenarioHistoryDB     ScenarioHistoryDB
	ParcelResultDB        ParcelResultDB
	ScenarioCollectionDBs []ScenarioCollectionDB
}
