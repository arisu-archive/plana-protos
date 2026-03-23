package protos

import (
	"github.com/arisu-archive/mapx"
)

type ConquestConquerWithBattleResultResponse struct {
	ResponsePacket
	ParcelResultDB               *ParcelResultDB `json:",omitempty,omitzero"`
	ConquestTileDB               *ConquestTileDB `json:",omitempty,omitzero"`
	ConquestInfoDB               *ConquestInfoDB `json:",omitempty,omitzero"`
	ConquestEventObjectDBWrapper []ConquestEventObjectDB
	DisplayInfos                 []*ConquestDisplayInfo
	StepAfterBattle              int32 `json:",omitempty,omitzero"`
	DisplayParcelByRewardTag     *mapx.OrderedMap[string, []*ParcelInfo]
}
