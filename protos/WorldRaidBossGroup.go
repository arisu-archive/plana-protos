package protos

import (
	"time"
)

type WorldRaidBossGroup struct {
	ContentsValueChangeDB
	GroupId int64 `json:",omitempty,omitzero"`
	BossSpawnTime time.Time `json:",omitempty,omitzero"`
	EliminateTime time.Time `json:",omitempty,omitzero"`
}
