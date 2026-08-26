package protos

type ScenarioListResponse struct {
	ResponsePacket
	ScenarioHistoryIds      []int64
	ScenarioGroupHistoryDBs []*ScenarioGroupHistoryDB
	ScenarioCollectionDBs   []*ScenarioCollectionDB
}
