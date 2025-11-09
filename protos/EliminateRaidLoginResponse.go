package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EliminateRaidLoginResponse struct {
	ResponsePacket
	SeasonType               flatdata.RaidSeasonType `json:",omitempty,omitzero"`
	CanReceiveRankingReward  bool                    `json:",omitempty,omitzero"`
	ReceiveLimitedRewardIds  []int64                 `json:",omitempty,omitzero"`
	SweepPointByRaidUniqueId map[int64]int64         `json:",omitempty,omitzero"`
	LastSettledRanking       int64                   `json:",omitempty,omitzero"`
	LastSettledTier          *int32                  `json:",omitempty,omitzero"`
}
