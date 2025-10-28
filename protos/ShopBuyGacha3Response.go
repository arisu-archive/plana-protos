package protos

type ShopBuyGacha3Response struct {
	ShopBuyGacha2Response
	Protocol Protocol `json:",omitempty,omitzero"`
	FreeRecruitHistoryDB ShopFreeRecruitHistoryDB `json:",omitempty,omitzero"`
	PickupFirstGetHistoryDBs []PickupFirstGetHistoryDB `json:",omitempty,omitzero"`
}
