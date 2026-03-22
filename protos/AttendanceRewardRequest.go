package protos

import (
	"github.com/arisu-archive/mapx"
)

type AttendanceRewardRequest struct {
	RequestPacket
	DayByBookUniqueId      *mapx.OrderedMap[int64, int64] `json:",omitempty,omitzero"`
	AttendanceBookUniqueId int64                          `json:",omitempty,omitzero"`
	Day                    int64                          `json:",omitempty,omitzero"`
}
