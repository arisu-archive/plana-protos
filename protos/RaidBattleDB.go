package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type RaidBattleDB struct {
	ContentType flatdata.ContentType `json:",omitempty,omitzero"`
	RaidUniqueId int64 `json:",omitempty,omitzero"`
	RaidBossIndex int32 `json:",omitempty,omitzero"`
	CurrentBossHP int64 `json:",omitempty,omitzero"`
	CurrentBossGroggy int64 `json:",omitempty,omitzero"`
	CurrentBossAIPhase int64 `json:",omitempty,omitzero"`
	BIEchelon string `json:",omitempty,omitzero"`
	IsClear bool `json:",omitempty,omitzero"`
	RaidMembers RaidMemberCollection `json:",omitempty,omitzero"`
	SubPartsHPs []int64 `json:",omitempty,omitzero"`
}
