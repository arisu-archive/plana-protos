package protos

import (
	"github.com/arisu-archive/mapx"
)

type EventContentCardShopListResponse struct {
	ResponsePacket
	CardShopElementDBs []CardShopElementDB                   `json:",omitempty,omitzero"`
	RewardHistory      *mapx.OrderedMap[int64, []ParcelInfo] `json:",omitempty,omitzero"`
}
