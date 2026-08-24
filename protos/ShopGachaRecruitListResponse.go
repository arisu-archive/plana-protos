package protos

type ShopGachaRecruitListResponse struct {
	ResponsePacket
	ShopRecruits              []*ShopRecruitDB
	ShopFreeRecruitHistoryDBs []*ShopFreeRecruitHistoryDB
	AccountLimitedGachaDBs    []*AccountLimitedGachaDB
}
