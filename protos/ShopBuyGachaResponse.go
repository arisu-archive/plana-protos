package protos

type ShopBuyGachaResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ConsumeResultDB ConsumeResultDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
