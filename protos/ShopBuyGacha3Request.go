package protos

type ShopBuyGacha3Request struct {
	ShopBuyGacha2Request
	Protocol Protocol `json:",omitempty,omitzero"`
	FreeRecruitId int64 `json:",omitempty,omitzero"`
	Cost ParcelCost `json:",omitempty,omitzero"`
}
