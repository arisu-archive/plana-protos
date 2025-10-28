package protos

type ShopBuyEligmaResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ConsumeResultDB ConsumeResultDB `json:",omitempty,omitzero"`
	ShopProductDB ShopProductDB `json:",omitempty,omitzero"`
}
