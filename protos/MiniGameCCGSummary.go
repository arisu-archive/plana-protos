package protos

import (
	"github.com/arisu-archive/mapx"
)

type MiniGameCCGSummary struct {
	RandomSeed        int32 `json:",omitempty,omitzero"`
	IsPlayerWin       bool  `json:",omitempty,omitzero"`
	Strikers          []MiniGameCCGCharacterDB
	TotalUsedCost     int32 `json:",omitempty,omitzero"`
	TotalDamageAmount int32 `json:",omitempty,omitzero"`
	TotalKillCount    int32 `json:",omitempty,omitzero"`
	TotalSkillCount   *mapx.OrderedMap[int64, int32]
	InputLogs         []string
}
