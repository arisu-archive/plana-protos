package protos

import (
	"github.com/arisu-archive/mapx"
)

type RoadPuzzleBoardSaveDB struct {
	UniqueId             int64                           `json:",omitempty,omitzero"`
	Round                int32                           `json:",omitempty,omitzero"`
	RecentRandomRailTile RoadPuzzleRailTileData          `json:",omitempty,omitzero"`
	RemainingTiles       *mapx.OrderedMap[string, int32] `json:",omitempty,omitzero"`
	PlacedRailTiles      []RoadPuzzleRailTileData        `json:",omitempty,omitzero"`
	RewardTiles          []RoadPuzzleTileRewardData      `json:",omitempty,omitzero"`
	IsTrainReadyToDepart bool                            `json:",omitempty,omitzero"`
}
