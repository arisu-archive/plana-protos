package protos

import (
	"github.com/arisu-archive/mapx"
)

type CharacterFavorGrowthRequest struct {
	RequestPacket
	TargetCharacterDBId       int64 `json:",omitempty,omitzero"`
	ConsumeItemDBIdsAndCounts *mapx.OrderedMap[int64, int32]
}
