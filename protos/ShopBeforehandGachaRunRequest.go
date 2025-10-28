package protos

type ShopBeforehandGachaRunRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ShopUniqueId int64 `json:",omitempty,omitzero"`
	GoodsId int64 `json:",omitempty,omitzero"`
}
