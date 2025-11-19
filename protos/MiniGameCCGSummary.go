package protos

type MiniGameCCGSummary struct {
	RandomSeed        int32                    `json:",omitempty,omitzero"`
	IsPlayerWin       bool                     `json:",omitempty,omitzero"`
	Strikers          []MiniGameCCGCharacterDB `json:",omitempty,omitzero"`
	TotalUsedCost     int32                    `json:",omitempty,omitzero"`
	TotalDamageAmount int32                    `json:",omitempty,omitzero"`
	TotalKillCount    int32                    `json:",omitempty,omitzero"`
	TotalSkillCount   map[int64]int32          `json:",omitempty,omitzero"`
	InputLogs         []string                 `json:",omitempty,omitzero"`
}
