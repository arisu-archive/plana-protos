package protos

import (
	"github.com/arisu-archive/mapx"
)

type AccountCurrencySyncResponse struct {
	ResponsePacket
	AccountCurrencyDB AccountCurrencyDB               `json:",omitempty,omitzero"`
	ExpiredCurrency   *mapx.OrderedMap[string, int64] `json:",omitempty,omitzero"`
}
