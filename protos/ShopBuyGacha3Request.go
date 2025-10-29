package protos

type ShopBuyGacha3Request struct {
	ShopBuyGacha2Request
	FreeRecruitId int64 `json:",omitempty,omitzero"`
	Cost ParcelCost `json:",omitempty,omitzero"`
}
