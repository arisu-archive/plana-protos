package protos

import (
	"time"
)

type AttendanceHistoryDB struct {
	ServerId int64 `json:",omitempty,omitzero"`
	AttendanceBookUniqueId int64 `json:",omitempty,omitzero"`
	AttendedDay map[int64]time.Time `json:",omitempty,omitzero"`
	Expired bool `json:",omitempty,omitzero"`
}
