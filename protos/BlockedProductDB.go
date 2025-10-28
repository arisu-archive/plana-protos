package protos

import (
	"time"
)

type BlockedProductDB struct {
	CashProductId int64 `json:",omitempty,omitzero"`
	MarketBlockType ShopCashBlockType `json:",omitempty,omitzero"`
	BeginDate time.Time `json:",omitempty,omitzero"`
	EndDate time.Time `json:",omitempty,omitzero"`
}
