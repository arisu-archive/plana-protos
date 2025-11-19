package protos

type ArenaPlayerInfoDB struct {
	CurrentSeasonId          int64  `json:",omitempty,omitzero"`
	PlayerGroupId            int64  `json:",omitempty,omitzero"`
	CurrentRank              int64  `json:",omitempty,omitzero"`
	SeasonRecord             int64  `json:",omitempty,omitzero"`
	AllTimeRecord            int64  `json:",omitempty,omitzero"`
	CumulativeTimeReward     int64  `json:",omitempty,omitzero"`
	TimeRewardLastUpdateTime MxTime `json:",omitempty,omitzero"`
	BattleEnterActiveTime    MxTime `json:",omitempty,omitzero"`
	DailyRewardActiveTime    MxTime `json:",omitempty,omitzero"`
}
