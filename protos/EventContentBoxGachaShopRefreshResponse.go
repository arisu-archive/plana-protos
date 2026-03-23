package protos

import (
	"github.com/arisu-archive/mapx"
)

type EventContentBoxGachaShopRefreshResponse struct {
	ResponsePacket
	BoxGachaDB             *EventContentBoxGachaDB `json:",omitempty,omitzero"`
	BoxGachaGroupIdByCount *mapx.OrderedMap[int64, int64]
}
