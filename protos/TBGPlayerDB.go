package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type TBGPlayerDB struct {
	Location HexLocation `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	HitPoint int32 `json:",omitempty,omitzero"`
	DiceId int64 `json:",omitempty,omitzero"`
	DiceProbModifyParams map[flatdata.TBGProbModifyCondition]int32 `json:",omitempty,omitzero"`
	Items []TBGItemDB `json:",omitempty,omitzero"`
	TemporaryItem TBGItemDB `json:",omitempty,omitzero"`
	ItemEffects []TBGItemEffectDB `json:",omitempty,omitzero"`
}
