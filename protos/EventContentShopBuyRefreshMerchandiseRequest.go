package protos

type EventContentShopBuyRefreshMerchandiseRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	ShopUniqueIds []int64 `json:",omitempty,omitzero"`
}
