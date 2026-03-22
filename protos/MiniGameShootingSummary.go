package protos

import (
	"github.com/arisu-archive/mapx"
)

type MiniGameShootingSummary struct {
	EventContentId    int64                          `json:",omitempty,omitzero"`
	StageId           int64                          `json:",omitempty,omitzero"`
	PlayerCharacterId int64                          `json:",omitempty,omitzero"`
	GeasIds           []int64                        `json:",omitempty,omitzero"`
	SectionCount      int64                          `json:",omitempty,omitzero"`
	ArriveSection     int64                          `json:",omitempty,omitzero"`
	LeftTimeSec       float32                        `json:",omitempty,omitzero"`
	ProgressedTimeSec float32                        `json:",omitempty,omitzero"`
	KillEnemies       *mapx.OrderedMap[int64, int32] `json:",omitempty,omitzero"`
	IsWin             bool                           `json:",omitempty,omitzero"`
}
