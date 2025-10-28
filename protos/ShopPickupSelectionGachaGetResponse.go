package protos

type ShopPickupSelectionGachaGetResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PickupCharacterSelection map[int64]int64 `json:",omitempty,omitzero"`
}
