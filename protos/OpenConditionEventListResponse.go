package protos

import (
	"github.com/arisu-archive/mapx"
)

type OpenConditionEventListResponse struct {
	ResponsePacket
	ConquestTiles         *mapx.OrderedMap[int64, []ConquestTileDB]       `json:",omitempty,omitzero"`
	WorldRaidLocalBossDBs *mapx.OrderedMap[int64, []WorldRaidLocalBossDB] `json:",omitempty,omitzero"`
}
