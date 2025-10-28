package protos

import (
	"time"
)

type MiniGameDreamMakerEndingDB struct {
	EventContentId int64 `json:",omitempty,omitzero"`
	EndingId int64 `json:",omitempty,omitzero"`
	ClearCount int64 `json:",omitempty,omitzero"`
	ClearDate time.Time `json:",omitempty,omitzero"`
}
