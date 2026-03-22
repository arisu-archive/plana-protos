package protos

import (
	"github.com/arisu-archive/mapx"
)

type EliminateRaidUserDB struct {
	RaidUserDB
	BossGroupToRankingPoint *mapx.OrderedMap[string, int64]
}
