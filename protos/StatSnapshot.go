package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type StatSnapshot struct {
	Stat  flatdata.StatType
	Start int64 `json:",omitempty,omitzero"`
	End   int64 `json:",omitempty,omitzero"`
}
