package protos

import (
	"github.com/arisu-archive/mapx"
)

type EventContentCardShopPurchaseAllResponse struct {
	ResponsePacket
	ParcelResultDB             *ParcelResultDB `json:",omitempty,omitzero"`
	CardShopElementDBs         []*CardShopElementDB
	CardShopPurchaseHistoryDBs []*CardShopPurchaseHistoryDB
	RewardHistory              *mapx.OrderedMap[int64, []*ParcelInfo]
}
