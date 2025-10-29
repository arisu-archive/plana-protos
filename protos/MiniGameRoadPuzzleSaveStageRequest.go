package protos

type MiniGameRoadPuzzleSaveStageRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	Round int64 `json:",omitempty,omitzero"`
	PlaceRailTiles []RoadPuzzleRailTileData `json:",omitempty,omitzero"`
}
