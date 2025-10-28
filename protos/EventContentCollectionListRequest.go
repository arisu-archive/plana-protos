package protos

type EventContentCollectionListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	GroupId *int64 `json:",omitempty,omitzero"`
}
