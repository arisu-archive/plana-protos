package protos

import (
	"github.com/arisu-archive/mapx"
)

type WorldRaidProgressDB struct {
	SeasonId          int64                           `json:",omitempty,omitzero"`
	PhaseId           int64                           `json:",omitempty,omitzero"`
	Maps              *mapx.OrderedMap[string, int64] `json:",omitempty,omitzero"`
	CarrierSkills     *mapx.OrderedMap[string, int32] `json:",omitempty,omitzero"`
	ClearConditionIds []int64                         `json:",omitempty,omitzero"`
}
