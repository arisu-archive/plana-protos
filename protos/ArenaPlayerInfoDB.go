package protos

import (
	"time"
)

type ArenaPlayerInfoDB struct {
	CurrentSeasonId int64 `json:",omitempty,omitzero"`
	PlayerGroupId int64 `json:",omitempty,omitzero"`
	CurrentRank int64 `json:",omitempty,omitzero"`
	SeasonRecord int64 `json:",omitempty,omitzero"`
	AllTimeRecord int64 `json:",omitempty,omitzero"`
	CumulativeTimeReward int64 `json:",omitempty,omitzero"`
	TimeRewardLastUpdateTime time.Time `json:",omitempty,omitzero"`
	BattleEnterActiveTime time.Time `json:",omitempty,omitzero"`
	DailyRewardActiveTime time.Time `json:",omitempty,omitzero"`
}
