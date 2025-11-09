package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ConquestMainStoryConquerWithBattleResultResponse struct {
	ResponsePacket
	ParcelResultDB           ParcelResultDB                      `json:",omitempty,omitzero"`
	ConquestTileDB           ConquestTileDB                      `json:",omitempty,omitzero"`
	ConquestInfoDB           ConquestInfoDB                      `json:",omitempty,omitzero"`
	DisplayInfos             []ConquestDisplayInfo               `json:",omitempty,omitzero"`
	StepAfterBattle          int32                               `json:",omitempty,omitzero"`
	DisplayParcelByRewardTag map[flatdata.RewardTag][]ParcelInfo `json:",omitempty,omitzero"`
}
