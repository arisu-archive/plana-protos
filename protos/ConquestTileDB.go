package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
	"time"
)

type ConquestTileDB struct {
	EventContentId int64 `json:",omitempty,omitzero"`
	Difficulty flatdata.StageDifficulty `json:",omitempty,omitzero"`
	TileUniqueId int64 `json:",omitempty,omitzero"`
	TileState flatdata.TileState `json:",omitempty,omitzero"`
	Level int64 `json:",omitempty,omitzero"`
	StarFlags []bool `json:",omitempty,omitzero"`
	CreateTime time.Time `json:",omitempty,omitzero"`
	IsThreeStarClear bool `json:",omitempty,omitzero"`
	IsAnyStarClear bool `json:",omitempty,omitzero"`
}
