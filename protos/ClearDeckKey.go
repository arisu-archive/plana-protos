package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ClearDeckKey struct {
	ContentType flatdata.ContentType
	Arguments   []int64
}
