package protos

type ShopBuyGacha3Response struct {
	ShopBuyGacha2Response
	FreeRecruitHistoryDB ShopFreeRecruitHistoryDB `json:",omitempty,omitzero"`
	PickupFirstGetHistoryDBs []PickupFirstGetHistoryDB `json:",omitempty,omitzero"`
}
