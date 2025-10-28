package protos

type ShopBuyGachaRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	GoodsId int64 `json:",omitempty,omitzero"`
	ShopUniqueId int64 `json:",omitempty,omitzero"`
}
