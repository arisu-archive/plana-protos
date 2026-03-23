package protos

type ShopBuyGacha2Response struct {
	ResponsePacket
	UpdateTime     MxTime `json:",omitempty,omitzero"`
	GemBonusRemain int64  `json:",omitempty,omitzero"`
	GemPaidRemain  int64  `json:",omitempty,omitzero"`
	ConsumedItems  []*ItemDB
	GachaResults   []*GachaResult
	AcquiredItems  []*ItemDB
}
