package protos

type EventContentFortuneGachaPurchaseResponse struct {
	ResponsePacket
	FortuneGachaShopUniqueId int64 `json:",omitempty,omitzero"`
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
