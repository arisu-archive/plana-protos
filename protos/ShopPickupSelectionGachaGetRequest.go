package protos

type ShopPickupSelectionGachaGetRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ShopRecruitId int64 `json:",omitempty,omitzero"`
}
