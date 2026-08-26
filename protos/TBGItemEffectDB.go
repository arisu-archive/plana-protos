package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type TBGItemEffectDB struct {
	ITBGItemEffectDB
	ItemUniqueId           int64                      `json:"id,omitempty,omitzero"`
	ItemType               flatdata.TBGItemType       `json:"item,omitempty,omitzero"`
	EffectType             flatdata.TBGItemEffectType `json:"eff,omitempty,omitzero"`
	RemainEncounterCounter int32                      `json:"rc,omitempty,omitzero"`
}
