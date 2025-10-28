package protos

type ShopBeforehandGachaPickResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	GachaResults []GachaResult `json:",omitempty,omitzero"`
	AcquiredItems []ItemDB `json:",omitempty,omitzero"`
}
