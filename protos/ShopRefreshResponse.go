package protos

type ShopRefreshResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ShopInfoDB     ShopInfoDB     `json:",omitempty,omitzero"`
}
