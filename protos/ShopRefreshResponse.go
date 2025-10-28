package protos

type ShopRefreshResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ShopInfoDB ShopInfoDB `json:",omitempty,omitzero"`
}
