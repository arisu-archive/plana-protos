package protos

type MultiFloorRaidDB struct {
	SeasonId          int64  `json:",omitempty,omitzero"`
	ClearedDifficulty int32  `json:",omitempty,omitzero"`
	LastClearDate     MxTime `json:",omitempty,omitzero"`
	RewardDifficulty  int32  `json:",omitempty,omitzero"`
	LastRewardDate    MxTime `json:",omitempty,omitzero"`
	ClearBattleFrame  int32  `json:",omitempty,omitzero"`
}
