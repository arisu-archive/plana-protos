package protos

type ShopPickupSelectionGachaBuyResponse struct {
	ShopBuyGacha2Response
	Protocol Protocol `json:",omitempty,omitzero"`
	FreeRecruitHistoryDB ShopFreeRecruitHistoryDB `json:",omitempty,omitzero"`
}
