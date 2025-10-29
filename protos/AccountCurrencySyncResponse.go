package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type AccountCurrencySyncResponse struct {
	ResponsePacket
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	ExpiredCurrency map[flatdata.CurrencyTypes]int64 `json:",omitempty,omitzero"`
}
