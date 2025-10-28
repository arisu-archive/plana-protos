package protos

type MiniGameCCGStagePlayDB struct {
	StageId int64 `json:",omitempty,omitzero"`
	EnemyGroupId int64 `json:",omitempty,omitzero"`
	IsClear bool `json:",omitempty,omitzero"`
	RewardDBs []MiniGameCCGStageRewardDB `json:",omitempty,omitzero"`
	RerollPoint int32 `json:",omitempty,omitzero"`
}
