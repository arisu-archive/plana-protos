package protos

type EventContentBoxGachaShopListRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
