package protos

type ShopBuyAPRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ShopUniqueId int64 `json:",omitempty,omitzero"`
	PurchaseCount int64 `json:",omitempty,omitzero"`
}
