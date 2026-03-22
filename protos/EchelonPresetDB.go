package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EchelonPresetDB struct {
	GroupIndex             int32 `json:",omitempty,omitzero"`
	Index                  int32 `json:",omitempty,omitzero"`
	Label                  string
	LeaderUniqueId         int64 `json:",omitempty,omitzero"`
	TSSInteractionUniqueId int64 `json:",omitempty,omitzero"`
	StrikerUniqueIds       []int64
	SpecialUniqueIds       []int64
	CombatStyleIndex       []int32
	MulliganUniqueIds      []int64
	ExtensionType          flatdata.EchelonExtensionType `json:",omitempty,omitzero"`
}
