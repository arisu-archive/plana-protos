package protos

type ShopPickupSelectionGachaBuyResponse struct {
	ShopBuyGacha2Response
	FreeRecruitHistoryDB ShopFreeRecruitHistoryDB `json:",omitempty,omitzero"`
}
