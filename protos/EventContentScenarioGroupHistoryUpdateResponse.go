package protos

type EventContentScenarioGroupHistoryUpdateResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ScenarioGroupHistoryDBs []ScenarioGroupHistoryDB `json:",omitempty,omitzero"`
	EventContentCollectionDBs []EventContentCollectionDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
