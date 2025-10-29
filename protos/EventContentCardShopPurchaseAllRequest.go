package protos

type EventContentCardShopPurchaseAllRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
}
