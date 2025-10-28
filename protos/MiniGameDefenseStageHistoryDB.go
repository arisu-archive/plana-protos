package protos

type MiniGameDefenseStageHistoryDB struct {
	StageId int64 `json:",omitempty,omitzero"`
	Star1Flag bool `json:",omitempty,omitzero"`
	Star2Flag bool `json:",omitempty,omitzero"`
	Star3Flag bool `json:",omitempty,omitzero"`
	FirstClearRewardReceive bool `json:",omitempty,omitzero"`
	StarRewardReceive bool `json:",omitempty,omitzero"`
}
