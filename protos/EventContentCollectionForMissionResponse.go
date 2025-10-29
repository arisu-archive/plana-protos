package protos

type EventContentCollectionForMissionResponse struct {
	ResponsePacket
	EventContentCollectionDBs []EventContentCollectionDB `json:",omitempty,omitzero"`
}
