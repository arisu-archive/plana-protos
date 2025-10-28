package protos

type ShopPickupSelectionGachaBuyRequest struct {
	ShopBuyGacha2Request
	Protocol Protocol `json:",omitempty,omitzero"`
	FreeRecruitId int64 `json:",omitempty,omitzero"`
	Cost ParcelCost `json:",omitempty,omitzero"`
}
