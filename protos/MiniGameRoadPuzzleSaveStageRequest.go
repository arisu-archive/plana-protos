package protos

type MiniGameRoadPuzzleSaveStageRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	Round int64 `json:",omitempty,omitzero"`
	PlaceRailTiles []RoadPuzzleRailTileData `json:",omitempty,omitzero"`
}
