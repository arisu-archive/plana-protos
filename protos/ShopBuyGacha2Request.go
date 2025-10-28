package protos

type ShopBuyGacha2Request struct {
	ShopBuyGachaRequest
	Protocol Protocol `json:",omitempty,omitzero"`
}
