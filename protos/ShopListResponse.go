package protos

type ShopListResponse struct {
	ResponsePacket
	ShopInfos []ShopInfoDB `json:",omitempty,omitzero"`
	ShopEligmaHistoryDBs []ShopEligmaHistoryDB `json:",omitempty,omitzero"`
}
