package protos

import (
	"github.com/arisu-archive/mapx"
)

type BeforehandGachaSnapshotDB struct {
	ShopUniqueId             int64 `json:",omitempty,omitzero"`
	GoodsId                  int64 `json:",omitempty,omitzero"`
	LastIndex                int64 `json:",omitempty,omitzero"`
	LastResults              []int64
	SavedIndex               *int64 `json:",omitempty,omitzero"`
	SavedResults             []int64
	SavedResultsBySavedIndex *mapx.OrderedMap[int32, []int64]
	PickedIndex              *int64 `json:",omitempty,omitzero"`
}
