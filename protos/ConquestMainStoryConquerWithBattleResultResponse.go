package protos

import (
	"github.com/arisu-archive/mapx"
)

type ConquestMainStoryConquerWithBattleResultResponse struct {
	ResponsePacket
	ParcelResultDB           ParcelResultDB
	ConquestTileDB           ConquestTileDB
	ConquestInfoDB           ConquestInfoDB
	DisplayInfos             []ConquestDisplayInfo
	StepAfterBattle          int32 `json:",omitempty,omitzero"`
	DisplayParcelByRewardTag *mapx.OrderedMap[string, []ParcelInfo]
}
