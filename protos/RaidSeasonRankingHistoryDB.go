package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
	"time"
)

type RaidSeasonRankingHistoryDB struct {
	ContentType flatdata.ContentType `json:",omitempty,omitzero"`
	AccountId int64 `json:",omitempty,omitzero"`
	SeasonId int64 `json:",omitempty,omitzero"`
	Ranking int64 `json:",omitempty,omitzero"`
	BestRankingPoint int64 `json:",omitempty,omitzero"`
	Tier int32 `json:",omitempty,omitzero"`
	ReceivedDate time.Time `json:",omitempty,omitzero"`
}
