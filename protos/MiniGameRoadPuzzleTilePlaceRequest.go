package protos

type MiniGameRoadPuzzleTilePlaceRequest struct {
	RequestPacket
	EventContentId  int64                  `json:",omitempty,omitzero"`
	UniqueId        int64                  `json:",omitempty,omitzero"`
	Round           int64                  `json:",omitempty,omitzero"`
	RailTileToPlace RoadPuzzleRailTileData `json:",omitempty,omitzero"`
}
