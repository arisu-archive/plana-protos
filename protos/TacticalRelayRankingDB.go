package protos

import (
	"time"
)

type TacticalRelayRankingDB struct {
	BestClearTime *time.Duration `json:",omitempty,omitzero"`
	Ranking       *BasisPoint    `json:",omitempty,omitzero"`
}
