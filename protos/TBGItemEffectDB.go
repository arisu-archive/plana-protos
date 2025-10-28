package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type TBGItemEffectDB struct {
	ITBGItemEffectDB
	ItemUniqueId int64 `json:",omitempty,omitzero"`
	ItemType flatdata.TBGItemType `json:",omitempty,omitzero"`
	EffectType flatdata.TBGItemEffectType `json:",omitempty,omitzero"`
	RemainEncounterCounter int32 `json:",omitempty,omitzero"`
}
