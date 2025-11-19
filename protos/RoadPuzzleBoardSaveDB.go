package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type RoadPuzzleBoardSaveDB struct {
	UniqueId             int64                                     `json:",omitempty,omitzero"`
	Round                int32                                     `json:",omitempty,omitzero"`
	RecentRandomRailTile RoadPuzzleRailTileData                    `json:",omitempty,omitzero"`
	RemainingTiles       map[flatdata.RoadPuzzleRailTileType]int32 `json:",omitempty,omitzero"`
	PlacedRailTiles      []RoadPuzzleRailTileData                  `json:",omitempty,omitzero"`
	RewardTiles          []RoadPuzzleTileRewardData                `json:",omitempty,omitzero"`
	IsTrainReadyToDepart bool                                      `json:",omitempty,omitzero"`
}
