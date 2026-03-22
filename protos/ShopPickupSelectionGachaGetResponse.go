package protos

import (
	"github.com/arisu-archive/mapx"
)

type ShopPickupSelectionGachaGetResponse struct {
	ResponsePacket
	PickupCharacterSelection *mapx.OrderedMap[int64, int64]
}
