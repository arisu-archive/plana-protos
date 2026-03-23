package protos

import (
	"github.com/arisu-archive/mapx"
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type RaidLobbyInfoDB struct {
	SeasonId                      int64   `json:",omitempty,omitzero"`
	Tier                          int32   `json:",omitempty,omitzero"`
	Ranking                       int64   `json:",omitempty,omitzero"`
	BestRankingPoint              int64   `json:",omitempty,omitzero"`
	TotalRankingPoint             int64   `json:",omitempty,omitzero"`
	ReceivedRankingRewardId       int64   `json:",omitempty,omitzero"`
	CanReceiveRankingReward       bool    `json:",omitempty,omitzero"`
	PlayingRaidDB                 *RaidDB `json:",omitempty,omitzero"`
	ReceiveRewardIds              []int64
	ReceiveLimitedRewardIds       []int64
	ParticipateCharacterServerIds []int64
	PlayableHighestDifficulty     *mapx.OrderedMap[string, flatdata.Difficulty]
	SweepPointByRaidUniqueId      *mapx.OrderedMap[int64, int64]
	SeasonStartDate               MxTime             `json:",omitempty,omitzero"`
	SeasonEndDate                 MxTime             `json:",omitempty,omitzero"`
	SettlementEndDate             MxTime             `json:",omitempty,omitzero"`
	NextSeasonId                  int64              `json:",omitempty,omitzero"`
	NextSeasonStartDate           MxTime             `json:",omitempty,omitzero"`
	NextSeasonEndDate             MxTime             `json:",omitempty,omitzero"`
	NextSettlementEndDate         MxTime             `json:",omitempty,omitzero"`
	ClanAssistUseInfo             *ClanAssistUseInfo `json:",omitempty,omitzero"`
	RemainFailCompensation        *mapx.OrderedMap[int32, bool]
}
