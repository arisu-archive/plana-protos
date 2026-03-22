package protos

type ShopListResponse struct {
	ResponsePacket
	ShopInfos            []ShopInfoDB
	ShopEligmaHistoryDBs []ShopEligmaHistoryDB
}
