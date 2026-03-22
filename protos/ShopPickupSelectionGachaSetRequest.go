package protos

import (
	"github.com/arisu-archive/mapx"
)

type ShopPickupSelectionGachaSetRequest struct {
	RequestPacket
	ShopRecruitID            int64 `json:",omitempty,omitzero"`
	PickupCharacterSelection *mapx.OrderedMap[int64, int64]
}
