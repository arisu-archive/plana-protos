package protos

type ShopBuyGachaResponse struct {
	ResponsePacket
	ConsumeResultDB ConsumeResultDB `json:",omitempty,omitzero"`
	ParcelResultDB  ParcelResultDB  `json:",omitempty,omitzero"`
}
