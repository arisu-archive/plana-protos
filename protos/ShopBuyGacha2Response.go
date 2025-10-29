package protos

import (
	"time"
)

type ShopBuyGacha2Response struct {
	ResponsePacket
	UpdateTime time.Time `json:",omitempty,omitzero"`
	GemBonusRemain int64 `json:",omitempty,omitzero"`
	GemPaidRemain int64 `json:",omitempty,omitzero"`
	ConsumedItems []ItemDB `json:",omitempty,omitzero"`
	GachaResults []GachaResult `json:",omitempty,omitzero"`
	AcquiredItems []ItemDB `json:",omitempty,omitzero"`
}
