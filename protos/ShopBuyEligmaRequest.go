package protos

type ShopBuyEligmaRequest struct {
	RequestPacket
	GoodsUniqueId     int64 `json:",omitempty,omitzero"`
	ShopUniqueId      int64 `json:",omitempty,omitzero"`
	CharacterUniqueId int64 `json:",omitempty,omitzero"`
	PurchaseCount     int64 `json:",omitempty,omitzero"`
}
