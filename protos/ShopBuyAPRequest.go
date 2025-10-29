package protos

type ShopBuyAPRequest struct {
	RequestPacket
	ShopUniqueId int64 `json:",omitempty,omitzero"`
	PurchaseCount int64 `json:",omitempty,omitzero"`
}
