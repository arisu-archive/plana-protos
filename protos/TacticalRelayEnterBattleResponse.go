package protos

import (
	"github.com/arisu-archive/mapx"
)

type TacticalRelayEnterBattleResponse struct {
	ResponsePacket
	ParcelResultDB                    *ParcelResultDB `json:",omitempty,omitzero"`
	AssistCharacterDBsByEchelonNumber *mapx.OrderedMap[int64, *AssistCharacterDB]
	AssistCharacterDBs                []*AssistCharacterDB
}
