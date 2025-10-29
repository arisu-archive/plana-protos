package protos

type ShopBuyRefreshMerchandiseRequest struct {
	RequestPacket
	ShopUniqueIds []int64 `json:",omitempty,omitzero"`
}
