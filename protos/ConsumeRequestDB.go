package protos

import (
	"github.com/arisu-archive/mapx"
)

type ConsumeRequestDB struct {
	ConsumeItemServerIdAndCounts      *mapx.OrderedMap[int64, int64] `json:",omitempty,omitzero"`
	ConsumeEquipmentServerIdAndCounts *mapx.OrderedMap[int64, int64] `json:",omitempty,omitzero"`
	ConsumeFurnitureServerIdAndCounts *mapx.OrderedMap[int64, int64] `json:",omitempty,omitzero"`
}
