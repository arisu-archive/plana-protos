package protos

type EventContentShopListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ShopInfos []ShopInfoDB `json:",omitempty,omitzero"`
	ShopEligmaHistoryDBs []ShopEligmaHistoryDB `json:",omitempty,omitzero"`
}
