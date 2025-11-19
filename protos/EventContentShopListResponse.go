package protos

type EventContentShopListResponse struct {
	ResponsePacket
	ShopInfos            []ShopInfoDB          `json:",omitempty,omitzero"`
	ShopEligmaHistoryDBs []ShopEligmaHistoryDB `json:",omitempty,omitzero"`
}
