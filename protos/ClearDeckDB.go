package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ClearDeckDB struct {
	ClearDeckCharacterDBs []ClearDeckCharacterDB `json:",omitempty,omitzero"`
	MulliganUniqueIds []int64 `json:",omitempty,omitzero"`
	LeaderUniqueId int64 `json:",omitempty,omitzero"`
	TSSInteractionUniqueId int64 `json:",omitempty,omitzero"`
	EchelonType flatdata.EchelonType `json:",omitempty,omitzero"`
	EchelonExtensionType int64 `json:",omitempty,omitzero"`
}
