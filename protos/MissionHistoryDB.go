package protos

import (
	"time"
)

type MissionHistoryDB struct {
	MissionUniqueId int64 `json:",omitempty,omitzero"`
	CompleteTime time.Time `json:",omitempty,omitzero"`
	Expired bool `json:",omitempty,omitzero"`
}
