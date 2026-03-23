package protos

import (
	"github.com/arisu-archive/mapx"
)

type TBGPlayerDB struct {
	Location             *HexLocation `json:",omitempty,omitzero"`
	EventContentId       int64        `json:",omitempty,omitzero"`
	HitPoint             int32        `json:",omitempty,omitzero"`
	DiceId               int64        `json:",omitempty,omitzero"`
	DiceProbModifyParams *mapx.OrderedMap[string, int32]
	Items                []*TBGItemDB
	TemporaryItem        *TBGItemDB `json:",omitempty,omitzero"`
	ItemEffects          []*TBGItemEffectDB
}
