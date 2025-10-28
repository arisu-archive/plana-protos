package protos

import (
	"time"
)

type PurchaseCountDB struct {
	ShopCashId int64 `json:",omitempty,omitzero"`
	PurchaseCount int32 `json:",omitempty,omitzero"`
	ResetDate time.Time `json:",omitempty,omitzero"`
	PurchaseDate *time.Time `json:",omitempty,omitzero"`
	ManualResetDate *time.Time `json:",omitempty,omitzero"`
}
