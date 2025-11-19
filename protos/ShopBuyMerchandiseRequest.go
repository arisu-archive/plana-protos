package protos

type ShopBuyMerchandiseRequest struct {
	RequestPacket
	IsRefreshGoods bool  `json:",omitempty,omitzero"`
	ShopUniqueId   int64 `json:",omitempty,omitzero"`
	GoodsId        int64 `json:",omitempty,omitzero"`
	PurchaseCount  int64 `json:",omitempty,omitzero"`
}
