package protos

import (
	"github.com/arisu-archive/mapx"
)

type SkipHistoryDB struct {
	Prologue int32 `json:",omitempty,omitzero"`
	Tutorial *mapx.OrderedMap[int32, int32]
}
