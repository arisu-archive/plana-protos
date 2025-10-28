package protos

type MiniGameRoadPuzzleTilePlaceResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RailTileToPlace RoadPuzzleRailTileData `json:",omitempty,omitzero"`
	SaveDB RoadPuzzleBoardSaveDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
