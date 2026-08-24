package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type RoadPuzzleRailTileData struct {
	RoadPuzzleTileData
	Type          flatdata.RoadPuzzleRailTileType `json:",omitempty,omitzero"`
	EntranceIndex int32                           `json:",omitempty,omitzero"`
	ExitIndex     int32                           `json:",omitempty,omitzero"`
	ResourcePath  string
}
