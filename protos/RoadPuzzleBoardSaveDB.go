package protos

type RoadPuzzleBoardSaveDB struct {
	UniqueId             int64                      `json:",omitempty,omitzero"`
	Round                int32                      `json:",omitempty,omitzero"`
	RecentRandomRailTile RoadPuzzleRailTileData     `json:",omitempty,omitzero"`
	RemainingTiles       map[string]int32           `json:",omitempty,omitzero"`
	PlacedRailTiles      []RoadPuzzleRailTileData   `json:",omitempty,omitzero"`
	RewardTiles          []RoadPuzzleTileRewardData `json:",omitempty,omitzero"`
	IsTrainReadyToDepart bool                       `json:",omitempty,omitzero"`
}
