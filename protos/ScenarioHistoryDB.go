package protos

import (
	"time"
)

type ScenarioHistoryDB struct {
	ScenarioUniqueId int64 `json:",omitempty,omitzero"`
	ClearDateTime time.Time `json:",omitempty,omitzero"`
}
