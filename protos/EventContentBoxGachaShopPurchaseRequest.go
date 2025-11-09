package protos

type EventContentBoxGachaShopPurchaseRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	PurchaseCount  int64 `json:",omitempty,omitzero"`
	PurchaseAll    bool  `json:",omitempty,omitzero"`
}
