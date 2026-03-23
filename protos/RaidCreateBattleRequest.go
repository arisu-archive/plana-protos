package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type RaidCreateBattleRequest struct {
	RequestPacket
	RaidUniqueId  int64 `json:",omitempty,omitzero"`
	IsPractice    bool  `json:",omitempty,omitzero"`
	Tags          []int32
	IsPublic      bool                `json:",omitempty,omitzero"`
	Difficulty    flatdata.Difficulty `json:",omitempty,omitzero"`
	AssistUseInfo *ClanAssistUseInfo  `json:",omitempty,omitzero"`
}
