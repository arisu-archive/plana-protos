package protos

type ScenarioListResponse struct {
	ResponsePacket
	ScenarioHistoryDBs      []*ScenarioHistoryDB
	ScenarioGroupHistoryDBs []*ScenarioGroupHistoryDB
	ScenarioCollectionDBs   []*ScenarioCollectionDB
}
