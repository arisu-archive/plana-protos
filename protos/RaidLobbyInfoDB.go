package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
	"time"
)

type RaidLobbyInfoDB struct {
	SeasonId int64 `json:",omitempty,omitzero"`
	Tier int32 `json:",omitempty,omitzero"`
	Ranking int64 `json:",omitempty,omitzero"`
	BestRankingPoint int64 `json:",omitempty,omitzero"`
	TotalRankingPoint int64 `json:",omitempty,omitzero"`
	ReceivedRankingRewardId int64 `json:",omitempty,omitzero"`
	CanReceiveRankingReward bool `json:",omitempty,omitzero"`
	PlayingRaidDB RaidDB `json:",omitempty,omitzero"`
	ReceiveRewardIds []int64 `json:",omitempty,omitzero"`
	ReceiveLimitedRewardIds []int64 `json:",omitempty,omitzero"`
	ParticipateCharacterServerIds []int64 `json:",omitempty,omitzero"`
	PlayableHighestDifficulty map[string]flatdata.Difficulty `json:",omitempty,omitzero"`
	SweepPointByRaidUniqueId map[int64]int64 `json:",omitempty,omitzero"`
	SeasonStartDate time.Time `json:",omitempty,omitzero"`
	SeasonEndDate time.Time `json:",omitempty,omitzero"`
	SettlementEndDate time.Time `json:",omitempty,omitzero"`
	NextSeasonId int64 `json:",omitempty,omitzero"`
	NextSeasonStartDate time.Time `json:",omitempty,omitzero"`
	NextSeasonEndDate time.Time `json:",omitempty,omitzero"`
	NextSettlementEndDate time.Time `json:",omitempty,omitzero"`
	ClanAssistUseInfo ClanAssistUseInfo `json:",omitempty,omitzero"`
	RemainFailCompensation map[int32]bool `json:",omitempty,omitzero"`
}
