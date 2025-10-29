package protos

type ShopBeforehandGachaRunRequest struct {
	RequestPacket
	ShopUniqueId int64 `json:",omitempty,omitzero"`
	GoodsId int64 `json:",omitempty,omitzero"`
}
