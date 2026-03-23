package protos

import (
	"github.com/arisu-archive/mapx"
)

type OpenConditionEventListResponse struct {
	ResponsePacket
	ConquestTiles         *mapx.OrderedMap[int64, []*ConquestTileDB]
	WorldRaidLocalBossDBs *mapx.OrderedMap[int64, []*WorldRaidLocalBossDB]
}
