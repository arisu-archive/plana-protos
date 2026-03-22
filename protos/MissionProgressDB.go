package protos

import (
	"github.com/arisu-archive/mapx"
)

type MissionProgressDB struct {
	MissionUniqueId    int64  `json:",omitempty,omitzero"`
	Complete           bool   `json:",omitempty,omitzero"`
	StartTime          MxTime `json:",omitempty,omitzero"`
	ProgressParameters *mapx.OrderedMap[int64, int64]
}
