package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type WorldRaidBossGroup struct {
	ContentsValueChangeDB
	ContentType   flatdata.ContentType `json:",omitempty,omitzero"`
	GroupId       int64                `json:",omitempty,omitzero"`
	BossSpawnTime MxTime               `json:",omitempty,omitzero"`
	EliminateTime MxTime               `json:",omitempty,omitzero"`
}
