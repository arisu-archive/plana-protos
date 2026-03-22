package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type WorldRaidWorldBossDB struct {
	ContentType  flatdata.ContentType
	GroupId      int64 `json:",omitempty,omitzero"`
	HP           int64 `json:",omitempty,omitzero"`
	Participants int64 `json:",omitempty,omitzero"`
}
