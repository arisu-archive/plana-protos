package protos

import (
	"github.com/arisu-archive/mapx"
)

type OpenConditionEventListRequest struct {
	RequestPacket
	ConquestEventIds           []int64                          `json:",omitempty,omitzero"`
	WorldRaidSeasonAndGroupIds *mapx.OrderedMap[int64, []int64] `json:",omitempty,omitzero"`
}
