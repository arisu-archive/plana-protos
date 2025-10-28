package protos

type EventContentBoxGachaShopPurchaseRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	PurchaseCount int64 `json:",omitempty,omitzero"`
	PurchaseAll bool `json:",omitempty,omitzero"`
}
