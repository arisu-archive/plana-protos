package protos

import (
	"github.com/arisu-archive/mapx"
)

type EventContentBoxGachaShopPurchaseResponse struct {
	ResponsePacket
	ParcelResultDB         ParcelResultDB                 `json:",omitempty,omitzero"`
	BoxGachaDB             EventContentBoxGachaDB         `json:",omitempty,omitzero"`
	BoxGachaGroupIdByCount *mapx.OrderedMap[int64, int64] `json:",omitempty,omitzero"`
	BoxGachaElements       []EventContentBoxGachaElement  `json:",omitempty,omitzero"`
}
