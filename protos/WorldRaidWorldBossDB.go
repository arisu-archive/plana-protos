package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type WorldRaidWorldBossDB struct {
	ContentType  flatdata.ContentType `json:",omitempty,omitzero"`
	GroupId      int64                `json:",omitempty,omitzero"`
	HP           int64                `json:",omitempty,omitzero"`
	Participants int64                `json:",omitempty,omitzero"`
}
