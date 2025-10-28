package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type TBGHexaMapDB struct {
	MapType flatdata.TBGThemaType `json:",omitempty,omitzero"`
	Objects map[int64]TBGHexaObjectDB `json:",omitempty,omitzero"`
	IsTutorial bool `json:",omitempty,omitzero"`
}
