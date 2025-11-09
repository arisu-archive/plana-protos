package protos

type MiniGameCCGSaveDB struct {
	CardDBIdProvider     int32                    `json:",omitempty,omitzero"`
	EventContentId       int64                    `json:",omitempty,omitzero"`
	CCGId                int64                    `json:",omitempty,omitzero"`
	GameOver             bool                     `json:",omitempty,omitzero"`
	Clear                bool                     `json:",omitempty,omitzero"`
	Strikers             []MiniGameCCGCharacterDB `json:",omitempty,omitzero"`
	Specials             []MiniGameCCGCharacterDB `json:",omitempty,omitzero"`
	OverflowedStrikerIds []int64                  `json:",omitempty,omitzero"`
	OverflowedSpecialIds []int64                  `json:",omitempty,omitzero"`
	Deck                 []MiniGameCCGCardDB      `json:",omitempty,omitzero"`
	LevelId              int64                    `json:",omitempty,omitzero"`
	CurrentNodeId        int64                    `json:",omitempty,omitzero"`
	RewardPoint          int32                    `json:",omitempty,omitzero"`
	CurrentStageIndex    int32                    `json:",omitempty,omitzero"`
	LastStageIndex       int32                    `json:",omitempty,omitzero"`
	CurrentStageDB       MiniGameCCGStagePlayDB   `json:",omitempty,omitzero"`
	ClearedHistoryDBs    []MiniGameCCGPlayHistory `json:",omitempty,omitzero"`
	Perks                []int64                  `json:",omitempty,omitzero"`
	TotalUsedCost        int32                    `json:",omitempty,omitzero"`
	TotalDamageAmount    int32                    `json:",omitempty,omitzero"`
	TotalKillCount       int32                    `json:",omitempty,omitzero"`
	TotalSkillCount      map[int64]int32          `json:",omitempty,omitzero"`
}
