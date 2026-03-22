package protos

import (
	"github.com/arisu-archive/mapx"
)

type ConsumeResultDB struct {
	RemovedItemServerIds                    []int64
	RemovedEquipmentServerIds               []int64
	RemovedFurnitureServerIds               []int64
	UsedItemServerIdAndRemainingCounts      *mapx.OrderedMap[int64, int64]
	UsedEquipmentServerIdAndRemainingCounts *mapx.OrderedMap[int64, int64]
	UsedFurnitureServerIdAndRemainingCounts *mapx.OrderedMap[int64, int64]
}
