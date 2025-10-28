package protos

type ShopBuyRefreshMerchandiseRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ShopUniqueIds []int64 `json:",omitempty,omitzero"`
}
