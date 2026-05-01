package protos

import (
	"encoding/json"
	"fmt"
)

type BattleSummary struct {
	HashKey            int64                       `json:",omitempty,omitzero"`
	IsBossBattle       bool                        `json:",omitempty,omitzero"`
	BattleType         BattleTypes                 `json:",omitempty,omitzero"`
	StageId            int64                       `json:",omitempty,omitzero"`
	GroundId           int64                       `json:",omitempty,omitzero"`
	Winner             BattleSummary_WinnerDefault `json:",omitzero"`
	EndType            BattleEndType               `json:",omitempty,omitzero"`
	EndFrame           int32                       `json:",omitempty,omitzero"`
	Group01Summary     *GroupSummary               `json:",omitempty,omitzero"`
	Group02Summary     *GroupSummary               `json:",omitempty,omitzero"`
	WeekDungeonSummary *WeekDungeonSummary         `json:",omitempty,omitzero"`
	RaidSummary        *RaidSummary                `json:",omitempty,omitzero"`
	TouchCountSummary  *ExcessiveTouch             `json:",omitempty,omitzero"`
	ArenaSummary       *ArenaSummary               `json:",omitempty,omitzero"`
	ContinueCount      int32                       `json:",omitempty,omitzero"`
	ElapsedRealtime    float32                     `json:",omitempty,omitzero"`
	IsAbort            bool                        `json:",omitempty,omitzero"`
	IsDefeatBattle     bool                        `json:",omitempty,omitzero"`
}

type BattleSummary_WinnerDefault string

func (s BattleSummary_WinnerDefault) IsZero() bool {
	return s == BattleSummary_WinnerDefault("None")
}

func (x *BattleSummary) UnmarshalJSON(data []byte) error {
	type aliasBattleSummary BattleSummary
	v := aliasBattleSummary{
		Winner: BattleSummary_WinnerDefault("None"),
	}
	if err := json.Unmarshal(data, &v); err != nil { //nolint:musttag // alias inherits json tags from the source struct via `type aliasT T`; linter walks the AST and misses the inheritance.
		return fmt.Errorf("unmarshal BattleSummary: %w", err)
	}
	*x = BattleSummary(v)
	return nil
}
