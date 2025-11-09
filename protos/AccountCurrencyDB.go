package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type AccountCurrencyDB struct {
	AccountLevel           int64                             `json:",omitempty,omitzero"`
	AcademyLocationRankSum int64                             `json:",omitempty,omitzero"`
	CurrencyDict           map[flatdata.CurrencyTypes]int64  `json:",omitempty,omitzero"`
	UpdateTimeDict         map[flatdata.CurrencyTypes]MxTime `json:",omitempty,omitzero"`
}
