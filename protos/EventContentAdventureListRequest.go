package protos

type EventContentAdventureListRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
