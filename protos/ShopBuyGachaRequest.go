package protos

type ShopBuyGachaRequest struct {
	RequestPacket
	GoodsId      int64 `json:",omitempty,omitzero"`
	ShopUniqueId int64 `json:",omitempty,omitzero"`
}
