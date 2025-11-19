package protos

type EventContentShopRefreshResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ShopInfoDB     ShopInfoDB     `json:",omitempty,omitzero"`
}
