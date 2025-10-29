package protos

type ShopBeforehandGachaPickResponse struct {
	ResponsePacket
	GachaResults []GachaResult `json:",omitempty,omitzero"`
	AcquiredItems []ItemDB `json:",omitempty,omitzero"`
}
