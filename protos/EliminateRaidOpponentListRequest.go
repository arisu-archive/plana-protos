package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EliminateRaidOpponentListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Rank *int64 `json:",omitempty,omitzero"`
	Score *int64 `json:",omitempty,omitzero"`
	BossGroupIndex *int32 `json:",omitempty,omitzero"`
	IsUpper bool `json:",omitempty,omitzero"`
	IsFirstRequest bool `json:",omitempty,omitzero"`
	SearchType flatdata.RankingSearchType `json:",omitempty,omitzero"`
}
