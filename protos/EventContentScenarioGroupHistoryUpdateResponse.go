package protos

type EventContentScenarioGroupHistoryUpdateResponse struct {
	ResponsePacket
	ScenarioGroupHistoryDBs []ScenarioGroupHistoryDB `json:",omitempty,omitzero"`
	EventContentCollectionDBs []EventContentCollectionDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
