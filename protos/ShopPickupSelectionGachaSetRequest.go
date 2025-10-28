package protos

type ShopPickupSelectionGachaSetRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ShopRecruitID int64 `json:",omitempty,omitzero"`
	PickupCharacterSelection map[int64]int64 `json:",omitempty,omitzero"`
}
