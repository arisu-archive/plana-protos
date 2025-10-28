package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EchelonPresetDB struct {
	GroupIndex int32 `json:",omitempty,omitzero"`
	Index int32 `json:",omitempty,omitzero"`
	Label string `json:",omitempty,omitzero"`
	LeaderUniqueId int64 `json:",omitempty,omitzero"`
	TSSInteractionUniqueId int64 `json:",omitempty,omitzero"`
	StrikerUniqueIds []int64 `json:",omitempty,omitzero"`
	SpecialUniqueIds []int64 `json:",omitempty,omitzero"`
	CombatStyleIndex []int32 `json:",omitempty,omitzero"`
	MulliganUniqueIds []int64 `json:",omitempty,omitzero"`
	ExtensionType flatdata.EchelonExtensionType `json:",omitempty,omitzero"`
}
