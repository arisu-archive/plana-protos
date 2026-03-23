package protos

import (
	"github.com/arisu-archive/mapx"
)

type EventContentCardShopListResponse struct {
	ResponsePacket
	CardShopElementDBs []*CardShopElementDB
	RewardHistory      *mapx.OrderedMap[int64, []*ParcelInfo]
}
