package protos

import (
	"github.com/arisu-archive/mapx"
)

type AccountCurrencySyncResponse struct {
	ResponsePacket
	AccountCurrencyDB AccountCurrencyDB
	ExpiredCurrency   *mapx.OrderedMap[string, int64]
}
