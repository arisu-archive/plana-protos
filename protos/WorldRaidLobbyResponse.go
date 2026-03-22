package protos

import (
	"github.com/arisu-archive/mapx"
)

type WorldRaidLobbyResponse struct {
	ResponsePacket
	ClearHistoryDBs     []WorldRaidClearHistoryDB
	LocalBossDBs        []WorldRaidLocalBossDB
	BossGroups          *mapx.OrderedMap[int64, []WorldRaidBossGroup]
	WorldRaidProgressDB WorldRaidProgressDB
}
