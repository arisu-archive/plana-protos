package protos

import (
	"time"
)

type CampaignStageHistoryDB struct {
	StoryUniqueId int64 `json:",omitempty,omitzero"`
	ChapterUniqueId int64 `json:",omitempty,omitzero"`
	StageUniqueId int64 `json:",omitempty,omitzero"`
	TacticClearCountWithRankSRecord int64 `json:",omitempty,omitzero"`
	ClearTurnRecord int64 `json:",omitempty,omitzero"`
	Star1Flag bool `json:",omitempty,omitzero"`
	Star2Flag bool `json:",omitempty,omitzero"`
	Star3Flag bool `json:",omitempty,omitzero"`
	LastPlay time.Time `json:",omitempty,omitzero"`
	TodayPlayCount int64 `json:",omitempty,omitzero"`
	TodayPurchasePlayCountHardStage int64 `json:",omitempty,omitzero"`
	FirstClearRewardReceive *time.Time `json:",omitempty,omitzero"`
	StarRewardReceive *time.Time `json:",omitempty,omitzero"`
	TodayPlayCountForUI int64 `json:",omitempty,omitzero"`
}
