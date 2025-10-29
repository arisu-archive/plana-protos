package protos

type EventContentCardShopListRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
