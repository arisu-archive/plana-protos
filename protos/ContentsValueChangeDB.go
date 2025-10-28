package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ContentsValueChangeDB struct {
	ContentsChangeType flatdata.ContentsChangeType `json:",omitempty,omitzero"`
}
