package protos

type ShopPickupSelectionGachaSetRequest struct {
	RequestPacket
	ShopRecruitID int64 `json:",omitempty,omitzero"`
	PickupCharacterSelection map[int64]int64 `json:",omitempty,omitzero"`
}
