package protos

import (
	"github.com/arisu-archive/mapx"
)

type AccountCurrencyDB struct {
	AccountLevel           int64 `json:",omitempty,omitzero"`
	AcademyLocationRankSum int64 `json:",omitempty,omitzero"`
	CurrencyDict           *mapx.OrderedMap[string, int64]
	UpdateTimeDict         *mapx.OrderedMap[string, MxTime]
}
