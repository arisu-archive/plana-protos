package protos

type EventContentBoxGachaShopRefreshRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
