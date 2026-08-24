package protos

type ShopGachaWishListSetRequest struct {
	RequestPacket
	ShopRecruitId      int64 `json:",omitempty,omitzero"`
	SelectCharacterIds []int64
}
