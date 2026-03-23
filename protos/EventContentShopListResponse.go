package protos

type EventContentShopListResponse struct {
	ResponsePacket
	ShopInfos            []*ShopInfoDB
	ShopEligmaHistoryDBs []*ShopEligmaHistoryDB
}
