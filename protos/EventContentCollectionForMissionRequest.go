package protos

type EventContentCollectionForMissionRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
