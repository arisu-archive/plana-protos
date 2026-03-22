package protos

import (
	"github.com/arisu-archive/mapx"
)

type RoadPuzzleBoardSaveDB struct {
	UniqueId             int64 `json:",omitempty,omitzero"`
	Round                int32 `json:",omitempty,omitzero"`
	RecentRandomRailTile RoadPuzzleRailTileData
	RemainingTiles       *mapx.OrderedMap[string, int32]
	PlacedRailTiles      []RoadPuzzleRailTileData
	RewardTiles          []RoadPuzzleTileRewardData
	IsTrainReadyToDepart bool `json:",omitempty,omitzero"`
}
