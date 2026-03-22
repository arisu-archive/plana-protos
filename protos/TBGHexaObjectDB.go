package protos

import (
	"github.com/arisu-archive/mapx"
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type TBGHexaObjectDB struct {
	ServerId                 int64                          `json:",omitempty,omitzero"`
	UniqueId                 int64                          `json:",omitempty,omitzero"`
	EncounterId              int64                          `json:",omitempty,omitzero"`
	MapType                  flatdata.TBGThemaType          `json:",omitempty,omitzero"`
	Location                 HexLocation                    `json:",omitempty,omitzero"`
	Activated                bool                           `json:",omitempty,omitzero"`
	HitPoint                 *int32                         `json:",omitempty,omitzero"`
	BeforeStoryOption        *int32                         `json:",omitempty,omitzero"`
	EncounterCostAlreadyPaid bool                           `json:",omitempty,omitzero"`
	IsFakeTreasure           *bool                          `json:",omitempty,omitzero"`
	FixRewardUniqueIdByIndex *mapx.OrderedMap[int32, int64] `json:",omitempty,omitzero"`
}
