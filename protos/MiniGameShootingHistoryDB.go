package protos

import (
	"time"
)

type MiniGameShootingHistoryDB struct {
	EventContentId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	ArriveSection int64 `json:",omitempty,omitzero"`
	LastUpdateDate time.Time `json:",omitempty,omitzero"`
	IsClearToday bool `json:",omitempty,omitzero"`
}
