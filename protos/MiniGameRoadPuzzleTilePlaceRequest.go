package protos

type MiniGameRoadPuzzleTilePlaceRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	Round int64 `json:",omitempty,omitzero"`
	RailTileToPlace RoadPuzzleRailTileData `json:",omitempty,omitzero"`
}
