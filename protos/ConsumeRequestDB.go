package protos

import (
	"github.com/arisu-archive/mapx"
)

type ConsumeRequestDB struct {
	ConsumeItemServerIdAndCounts      *mapx.OrderedMap[int64, int64]
	ConsumeEquipmentServerIdAndCounts *mapx.OrderedMap[int64, int64]
	ConsumeFurnitureServerIdAndCounts *mapx.OrderedMap[int64, int64]
}
