package protos

import (
	"time"
)

type AcademyDB struct {
	LastUpdate time.Time `json:",omitempty,omitzero"`
	ZoneVisitCharacterDBs map[int64][]VisitingCharacterDB `json:",omitempty,omitzero"`
	ZoneScheduleGroupRecords map[int64][]int64 `json:",omitempty,omitzero"`
}
