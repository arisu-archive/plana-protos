package protos

import (
	"github.com/arisu-archive/mapx"
)

type WorldRaidProgressDB struct {
	SeasonId          int64 `json:",omitempty,omitzero"`
	PhaseId           int64 `json:",omitempty,omitzero"`
	Maps              *mapx.OrderedMap[string, int64]
	CarrierSkills     *mapx.OrderedMap[string, int32]
	ClearConditionIds []int64
}
