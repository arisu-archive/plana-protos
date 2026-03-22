package protos

import (
	"github.com/arisu-archive/mapx"
)

type ConquestConquerWithBattleResultResponse struct {
	ResponsePacket
	ParcelResultDB               ParcelResultDB
	ConquestTileDB               ConquestTileDB
	ConquestInfoDB               ConquestInfoDB
	ConquestEventObjectDBWrapper []ConquestEventObjectDB
	DisplayInfos                 []ConquestDisplayInfo
	StepAfterBattle              int32 `json:",omitempty,omitzero"`
	DisplayParcelByRewardTag     *mapx.OrderedMap[string, []ParcelInfo]
}
