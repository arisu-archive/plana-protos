package protos

import (
	"github.com/arisu-archive/mapx"
)

type EventContentLocationDB struct {
	LocationId            int64 `json:",omitempty,omitzero"`
	Rank                  int64 `json:",omitempty,omitzero"`
	Exp                   int64 `json:",omitempty,omitzero"`
	ScheduleCount         int64 `json:",omitempty,omitzero"`
	ZoneVisitCharacterDBs *mapx.OrderedMap[int64, []VisitingCharacterDB]
}
