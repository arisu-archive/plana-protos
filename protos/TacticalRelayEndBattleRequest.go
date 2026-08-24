package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type TacticalRelayEndBattleRequest struct {
	RequestPacket
	SeasonId           int64          `json:",omitempty,omitzero"`
	StageId            int64          `json:",omitempty,omitzero"`
	EchelonId          int64          `json:",omitempty,omitzero"`
	Summary            *BattleSummary `json:",omitempty,omitzero"`
	AssistUseInfos     []*ClanAssistUseInfo
	SelectedRewardType flatdata.EngraveTreeType `json:",omitempty,omitzero"`
}
