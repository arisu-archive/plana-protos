package protos

type ShopBuyGacha3Response struct {
	ShopBuyGacha2Response
	AccountCurrencyDB            *AccountCurrencyDB        `json:",omitempty,omitzero"`
	FreeRecruitHistoryDB         *ShopFreeRecruitHistoryDB `json:",omitempty,omitzero"`
	PickupFirstGetHistoryDBs     []*PickupFirstGetHistoryDB
	ShopRecruitMileageHistoryDB  *ShopRecruitMileageHistoryDB `json:",omitempty,omitzero"`
	ShopRecruitMileageRewardList []*ParcelInfo
	AccountLimitedFlashSaleDBs   []*AccountLimitedFlashSaleDB
}
