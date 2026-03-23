package protos

type EventContentScenarioGroupHistoryUpdateResponse struct {
	ResponsePacket
	ScenarioGroupHistoryDBs   []ScenarioGroupHistoryDB
	EventContentCollectionDBs []EventContentCollectionDB
	ParcelResultDB            *ParcelResultDB `json:",omitempty,omitzero"`
}
