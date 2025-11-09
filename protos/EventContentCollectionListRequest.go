package protos

type EventContentCollectionListRequest struct {
	RequestPacket
	EventContentId int64  `json:",omitempty,omitzero"`
	GroupId        *int64 `json:",omitempty,omitzero"`
}
