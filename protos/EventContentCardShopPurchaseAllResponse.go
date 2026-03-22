package protos

import (
	"github.com/arisu-archive/mapx"
)

type EventContentCardShopPurchaseAllResponse struct {
	ResponsePacket
	ParcelResultDB             ParcelResultDB                        `json:",omitempty,omitzero"`
	CardShopElementDBs         []CardShopElementDB                   `json:",omitempty,omitzero"`
	CardShopPurchaseHistoryDBs []CardShopPurchaseHistoryDB           `json:",omitempty,omitzero"`
	RewardHistory              *mapx.OrderedMap[int64, []ParcelInfo] `json:",omitempty,omitzero"`
}
