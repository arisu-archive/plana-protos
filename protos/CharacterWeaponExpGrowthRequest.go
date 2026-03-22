package protos

import (
	"github.com/arisu-archive/mapx"
)

type CharacterWeaponExpGrowthRequest struct {
	RequestPacket
	TargetCharacterServerId  int64 `json:",omitempty,omitzero"`
	ConsumeUniqueIdAndCounts *mapx.OrderedMap[int64, int64]
}
