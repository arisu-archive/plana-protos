package protos

import (
	"github.com/arisu-archive/mapx"
)

type AccountCurrencyDB struct {
	AccountLevel           int64                            `json:",omitempty,omitzero"`
	AcademyLocationRankSum int64                            `json:",omitempty,omitzero"`
	CurrencyDict           *mapx.OrderedMap[string, int64]  `json:",omitempty,omitzero"`
	UpdateTimeDict         *mapx.OrderedMap[string, MxTime] `json:",omitempty,omitzero"`
}
