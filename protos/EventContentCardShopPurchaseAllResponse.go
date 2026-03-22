package protos

import (
	"github.com/arisu-archive/mapx"
)

type EventContentCardShopPurchaseAllResponse struct {
	ResponsePacket
	ParcelResultDB             ParcelResultDB
	CardShopElementDBs         []CardShopElementDB
	CardShopPurchaseHistoryDBs []CardShopPurchaseHistoryDB
	RewardHistory              *mapx.OrderedMap[int64, []ParcelInfo]
}
