package protos

type EventContentCollectionForMissionResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentCollectionDBs []EventContentCollectionDB `json:",omitempty,omitzero"`
}
