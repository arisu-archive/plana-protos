package protos

import (
	"github.com/arisu-archive/mapx"
)

type EliminateRaidLobbyInfoDB struct {
	RaidLobbyInfoDB
	OpenedBossGroups             []string                        `json:",omitempty,omitzero"`
	BestRankingPointPerBossGroup *mapx.OrderedMap[string, int64] `json:",omitempty,omitzero"`
}
