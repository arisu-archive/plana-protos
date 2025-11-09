package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type RaidLoginResponse struct {
	ResponsePacket
	SeasonType              flatdata.RaidSeasonType `json:",omitempty,omitzero"`
	CanReceiveRankingReward bool                    `json:",omitempty,omitzero"`
	LastSettledRanking      int64                   `json:",omitempty,omitzero"`
	LastSettledTier         *int32                  `json:",omitempty,omitzero"`
}
