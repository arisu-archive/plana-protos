package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type RaidBossDB struct {
	ContentType     flatdata.ContentType `json:",omitempty,omitzero"`
	BossIndex       int32                `json:",omitempty,omitzero"`
	BossCurrentHP   int64                `json:",omitempty,omitzero"`
	BossGroggyPoint int64                `json:",omitempty,omitzero"`
}
