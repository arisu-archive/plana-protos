package protos

type EventContentShopBuyMerchandiseRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	IsRefreshMerchandise bool `json:",omitempty,omitzero"`
	ShopUniqueId int64 `json:",omitempty,omitzero"`
	GoodsUniqueId int64 `json:",omitempty,omitzero"`
	PurchaseCount int64 `json:",omitempty,omitzero"`
}
