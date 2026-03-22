package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type WeekDungeonSummary struct {
	DungeonType flatdata.WeekDungeonType `json:",omitempty,omitzero"`
	FindGifts   []FindGiftSummary
}
