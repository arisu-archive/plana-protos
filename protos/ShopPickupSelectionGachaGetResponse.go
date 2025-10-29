package protos

type ShopPickupSelectionGachaGetResponse struct {
	ResponsePacket
	PickupCharacterSelection map[int64]int64 `json:",omitempty,omitzero"`
}
