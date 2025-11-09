package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type MonthlyProductPurchaseDB struct {
	ProductId           int64                   `json:",omitempty,omitzero"`
	PurchaseDate        MxTime                  `json:",omitempty,omitzero"`
	LastDailyRewardDate *MxTime                 `json:",omitempty,omitzero"`
	RewardEndDate       *MxTime                 `json:",omitempty,omitzero"`
	ProductTagType      flatdata.ProductTagType `json:",omitempty,omitzero"`
}
