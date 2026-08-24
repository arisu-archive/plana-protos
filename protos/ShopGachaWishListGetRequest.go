package protos

type ShopGachaWishListGetRequest struct {
	RequestPacket
	ShopRecruitId int64 `json:",omitempty,omitzero"`
}
