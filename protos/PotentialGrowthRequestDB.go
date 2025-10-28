package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type PotentialGrowthRequestDB struct {
	Type flatdata.PotentialStatBonusRateType `json:",omitempty,omitzero"`
	Level int32 `json:",omitempty,omitzero"`
}
