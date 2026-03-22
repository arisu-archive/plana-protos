package protos

import (
	"github.com/arisu-archive/mapx"
)

type AcademyDB struct {
	LastUpdate               MxTime                                         `json:",omitempty,omitzero"`
	ZoneVisitCharacterDBs    *mapx.OrderedMap[int64, []VisitingCharacterDB] `json:",omitempty,omitzero"`
	ZoneScheduleGroupRecords *mapx.OrderedMap[int64, []int64]               `json:",omitempty,omitzero"`
}
