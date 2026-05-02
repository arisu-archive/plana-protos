package protos

import (
	"github.com/arisu-archive/mapx"
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type TBGHexaMapDB struct {
	MapType    flatdata.TBGThemaType                     `json:",omitempty,omitzero"`
	Objects    *mapx.OrderedMap[int64, *TBGHexaObjectDB] `json:"obj"`
	IsTutorial bool                                      `json:"tutorial,omitempty,omitzero"`
}
