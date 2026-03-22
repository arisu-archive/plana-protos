package protos

type ShopBuyGacha3Response struct {
	ShopBuyGacha2Response
	FreeRecruitHistoryDB     ShopFreeRecruitHistoryDB
	PickupFirstGetHistoryDBs []PickupFirstGetHistoryDB
}
