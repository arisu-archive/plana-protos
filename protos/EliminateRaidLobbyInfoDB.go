package protos

import (
	"github.com/arisu-archive/mapx"
)

type EliminateRaidLobbyInfoDB struct {
	RaidLobbyInfoDB
	OpenedBossGroups             []string
	BestRankingPointPerBossGroup *mapx.OrderedMap[string, int64]
}
