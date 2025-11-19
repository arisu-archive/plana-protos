package protos

type EventContentShopBuyMerchandiseRequest struct {
	RequestPacket
	EventContentId       int64 `json:",omitempty,omitzero"`
	IsRefreshMerchandise bool  `json:",omitempty,omitzero"`
	ShopUniqueId         int64 `json:",omitempty,omitzero"`
	GoodsUniqueId        int64 `json:",omitempty,omitzero"`
	PurchaseCount        int64 `json:",omitempty,omitzero"`
}
