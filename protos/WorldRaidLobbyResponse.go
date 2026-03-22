package protos

import (
	"github.com/arisu-archive/mapx"
)

type WorldRaidLobbyResponse struct {
	ResponsePacket
	ClearHistoryDBs     []WorldRaidClearHistoryDB                     `json:",omitempty,omitzero"`
	LocalBossDBs        []WorldRaidLocalBossDB                        `json:",omitempty,omitzero"`
	BossGroups          *mapx.OrderedMap[int64, []WorldRaidBossGroup] `json:",omitempty,omitzero"`
	WorldRaidProgressDB WorldRaidProgressDB                           `json:",omitempty,omitzero"`
}
