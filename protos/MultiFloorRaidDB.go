package protos

import (
	"time"
)

type MultiFloorRaidDB struct {
	SeasonId int64 `json:",omitempty,omitzero"`
	ClearedDifficulty int32 `json:",omitempty,omitzero"`
	LastClearDate time.Time `json:",omitempty,omitzero"`
	RewardDifficulty int32 `json:",omitempty,omitzero"`
	LastRewardDate time.Time `json:",omitempty,omitzero"`
	ClearBattleFrame int32 `json:",omitempty,omitzero"`
}
