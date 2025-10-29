package protos

type MiniGameRoadPuzzleTilePlaceResponse struct {
	ResponsePacket
	RailTileToPlace RoadPuzzleRailTileData `json:",omitempty,omitzero"`
	SaveDB RoadPuzzleBoardSaveDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
