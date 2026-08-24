package protos

import (
	"github.com/arisu-archive/mapx"
)

type TacticalRelayEnterBattleRequest struct {
	RequestPacket
	SeasonId                      int64 `json:",omitempty,omitzero"`
	StageId                       int64 `json:",omitempty,omitzero"`
	AssistUseInfosByEchelonNumber *mapx.OrderedMap[int64, *ClanAssistUseInfo]
	AssistUseInfos                []*ClanAssistUseInfo
}
