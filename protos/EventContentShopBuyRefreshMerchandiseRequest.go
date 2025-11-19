package protos

type EventContentShopBuyRefreshMerchandiseRequest struct {
	RequestPacket
	EventContentId int64   `json:",omitempty,omitzero"`
	ShopUniqueIds  []int64 `json:",omitempty,omitzero"`
}
