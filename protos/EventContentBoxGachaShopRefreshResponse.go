package protos

import (
	"github.com/arisu-archive/mapx"
)

type EventContentBoxGachaShopRefreshResponse struct {
	ResponsePacket
	BoxGachaDB             EventContentBoxGachaDB
	BoxGachaGroupIdByCount *mapx.OrderedMap[int64, int64]
}
