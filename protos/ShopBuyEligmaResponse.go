package protos

type ShopBuyEligmaResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	ConsumeResultDB ConsumeResultDB `json:",omitempty,omitzero"`
	ShopProductDB ShopProductDB `json:",omitempty,omitzero"`
}
