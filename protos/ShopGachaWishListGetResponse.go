package protos

type ShopGachaWishListGetResponse struct {
	ResponsePacket
	AccountGachaWishDB *AccountGachaWishDB `json:",omitempty,omitzero"`
}
