package protos

import (
	"github.com/arisu-archive/mapx"
)

type BattleSummary struct {
	MulligansByEchelonIndex *mapx.OrderedMap[int32, []int64]
	HashKey                 int64                 `json:",omitempty,omitzero"`
	IsBossBattle            bool                  `json:",omitempty,omitzero"`
	BattleType              BattleTypes           `json:",omitempty,omitzero"`
	StageId                 int64                 `json:",omitempty,omitzero"`
	GroundId                int64                 `json:",omitempty,omitzero"`
	Winner                  string                `json:",omitempty,omitzero"`
	EndType                 BattleEndType         `json:",omitempty,omitzero"`
	EndFrame                int32                 `json:",omitempty,omitzero"`
	Group01Summary          *GroupSummary         `json:",omitempty,omitzero"`
	Group02Summary          *GroupSummary         `json:",omitempty,omitzero"`
	WeekDungeonSummary      *WeekDungeonSummary   `json:",omitempty,omitzero"`
	RaidSummary             *RaidSummary          `json:",omitempty,omitzero"`
	TouchCountSummary       *ExcessiveTouch       `json:",omitempty,omitzero"`
	ArenaSummary            *ArenaSummary         `json:",omitempty,omitzero"`
	TacticalRelaySummary    *TacticalRelaySummary `json:",omitempty,omitzero"`
	ContinueCount           int32                 `json:",omitempty,omitzero"`
	ElapsedRealtime         float32               `json:",omitempty,omitzero"`
	IsAbort                 bool                  `json:",omitempty,omitzero"`
	IsDefeatBattle          bool                  `json:",omitempty,omitzero"`
}
