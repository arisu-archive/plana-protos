package protos

type MiniGameCCGStageRewardDB struct {
	Type MiniGameCCGStageRewardType `json:",omitempty,omitzero"`
	RewardIndex int32 `json:",omitempty,omitzero"`
	RewardIds []int64 `json:",omitempty,omitzero"`
}
