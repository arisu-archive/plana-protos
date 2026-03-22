package protos

import (
	"github.com/arisu-archive/mapx"
)

type ConsumeResultDB struct {
	RemovedItemServerIds                    []int64                        `json:",omitempty,omitzero"`
	RemovedEquipmentServerIds               []int64                        `json:",omitempty,omitzero"`
	RemovedFurnitureServerIds               []int64                        `json:",omitempty,omitzero"`
	UsedItemServerIdAndRemainingCounts      *mapx.OrderedMap[int64, int64] `json:",omitempty,omitzero"`
	UsedEquipmentServerIdAndRemainingCounts *mapx.OrderedMap[int64, int64] `json:",omitempty,omitzero"`
	UsedFurnitureServerIdAndRemainingCounts *mapx.OrderedMap[int64, int64] `json:",omitempty,omitzero"`
}
