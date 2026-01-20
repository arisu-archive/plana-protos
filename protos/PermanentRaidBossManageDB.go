package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type PermanentRaidBossManageDB struct {
	GroupType     flatdata.RaidBossGroupType `json:",omitempty,omitzero"`
	LockStartDate MxTime                     `json:",omitempty,omitzero"`
	LockEndDate   MxTime                     `json:",omitempty,omitzero"`
}
