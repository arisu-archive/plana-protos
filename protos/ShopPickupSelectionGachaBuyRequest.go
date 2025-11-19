package protos

type ShopPickupSelectionGachaBuyRequest struct {
	ShopBuyGacha2Request
	FreeRecruitId int64      `json:",omitempty,omitzero"`
	Cost          ParcelCost `json:",omitempty,omitzero"`
}
