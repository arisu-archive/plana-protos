package protos

type ScenarioCollectionResponse struct {
	ResponsePacket
	ScenarioCollectionDBs []*ScenarioCollectionDB
}
