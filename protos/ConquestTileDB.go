package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ConquestTileDB struct {
	EventContentId   int64                    `json:",omitempty,omitzero"`
	Difficulty       flatdata.StageDifficulty `json:",omitempty,omitzero"`
	TileUniqueId     int64                    `json:",omitempty,omitzero"`
	TileState        flatdata.TileState
	Level            int64 `json:",omitempty,omitzero"`
	StarFlags        []bool
	CreateTime       MxTime `json:",omitempty,omitzero"`
	IsThreeStarClear bool   `json:",omitempty,omitzero"`
	IsAnyStarClear   bool   `json:",omitempty,omitzero"`
}
