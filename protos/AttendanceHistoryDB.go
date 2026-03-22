package protos

import (
	"github.com/arisu-archive/mapx"
)

type AttendanceHistoryDB struct {
	ServerId               int64                           `json:",omitempty,omitzero"`
	AttendanceBookUniqueId int64                           `json:",omitempty,omitzero"`
	AttendedDay            *mapx.OrderedMap[int64, MxTime] `json:",omitempty,omitzero"`
	Expired                bool                            `json:",omitempty,omitzero"`
}
