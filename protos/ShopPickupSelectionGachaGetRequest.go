package protos

type ShopPickupSelectionGachaGetRequest struct {
	RequestPacket
	ShopRecruitId int64 `json:",omitempty,omitzero"`
}
