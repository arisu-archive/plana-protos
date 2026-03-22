package protos

import (
	"github.com/arisu-archive/mapx"
)

type AcademyDB struct {
	LastUpdate               MxTime `json:",omitempty,omitzero"`
	ZoneVisitCharacterDBs    *mapx.OrderedMap[int64, []VisitingCharacterDB]
	ZoneScheduleGroupRecords *mapx.OrderedMap[int64, []int64]
}
