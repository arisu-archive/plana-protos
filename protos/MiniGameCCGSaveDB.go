package protos

import (
	"github.com/arisu-archive/mapx"
)

type MiniGameCCGSaveDB struct {
	CardDBIdProvider     int32 `json:",omitempty,omitzero"`
	EventContentId       int64 `json:",omitempty,omitzero"`
	CCGId                int64 `json:",omitempty,omitzero"`
	GameOver             bool  `json:",omitempty,omitzero"`
	Clear                bool  `json:",omitempty,omitzero"`
	Strikers             []*MiniGameCCGCharacterDB
	Specials             []*MiniGameCCGCharacterDB
	OverflowedStrikerIds []int64
	OverflowedSpecialIds []int64
	Deck                 []*MiniGameCCGCardDB
	LevelId              int64                   `json:",omitempty,omitzero"`
	CurrentNodeId        int64                   `json:",omitempty,omitzero"`
	RewardPoint          int32                   `json:",omitempty,omitzero"`
	CurrentStageIndex    int32                   `json:",omitempty,omitzero"`
	LastStageIndex       int32                   `json:",omitempty,omitzero"`
	CurrentStageDB       *MiniGameCCGStagePlayDB `json:",omitempty,omitzero"`
	ClearedHistoryDBs    []*MiniGameCCGPlayHistory
	Perks                []int64
	TotalUsedCost        int32 `json:",omitempty,omitzero"`
	TotalDamageAmount    int32 `json:",omitempty,omitzero"`
	TotalKillCount       int32 `json:",omitempty,omitzero"`
	TotalSkillCount      *mapx.OrderedMap[int64, int32]
}
