package protos

import (
	"github.com/arisu-archive/mapx"
)

type EventContentBoxGachaShopListResponse struct {
	ResponsePacket
	BoxGachaDB             EventContentBoxGachaDB
	BoxGachaGroupIdByCount *mapx.OrderedMap[int64, int64]
}
