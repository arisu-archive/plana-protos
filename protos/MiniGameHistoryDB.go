package protos

import (
	"time"
)

type MiniGameHistoryDB struct {
	EventContentId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	HighScore int64 `json:",omitempty,omitzero"`
	AccumulatedScore int64 `json:",omitempty,omitzero"`
	ClearDate time.Time `json:",omitempty,omitzero"`
	IsFullCombo bool `json:",omitempty,omitzero"`
}
