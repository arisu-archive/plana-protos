package protos

import (
	"github.com/arisu-archive/mapx"
)

type EventContentBoxGachaShopPurchaseResponse struct {
	ResponsePacket
	ParcelResultDB         ParcelResultDB
	BoxGachaDB             EventContentBoxGachaDB
	BoxGachaGroupIdByCount *mapx.OrderedMap[int64, int64]
	BoxGachaElements       []EventContentBoxGachaElement
}
