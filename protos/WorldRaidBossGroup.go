package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
	"time"
)

type WorldRaidBossGroup struct {
	ContentsValueChangeDB
	ContentsChangeType flatdata.ContentsChangeType `json:",omitempty,omitzero"`
	GroupId int64 `json:",omitempty,omitzero"`
	BossSpawnTime time.Time `json:",omitempty,omitzero"`
	EliminateTime time.Time `json:",omitempty,omitzero"`
}
