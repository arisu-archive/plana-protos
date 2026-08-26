package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ITBGItemInfo struct {
	UniqueId           int64                      `json:",omitempty,omitzero"`
	ItemType           flatdata.TBGItemType       `json:",omitempty,omitzero"`
	ItemEffectType     flatdata.TBGItemEffectType `json:",omitempty,omitzero"`
	ItemParameter      int32                      `json:",omitempty,omitzero"`
	EncounterCount     int32                      `json:",omitempty,omitzero"`
	LocalizeEtcId      string                     `json:",omitempty,omitzero"`
	Icon               string                     `json:",omitempty,omitzero"`
	BuffIcon           string                     `json:",omitempty,omitzero"`
	DiceEffectAniClip  string                     `json:",omitempty,omitzero"`
	BuffIconHUDVisible bool                       `json:",omitempty,omitzero"`
}
