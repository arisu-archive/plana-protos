package protos

type ShopListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ShopInfos []ShopInfoDB `json:",omitempty,omitzero"`
	ShopEligmaHistoryDBs []ShopEligmaHistoryDB `json:",omitempty,omitzero"`
}
