package protos

type MiniGameRoadPuzzleTilePlaceResponse struct {
	ResponsePacket
	RailTileToPlace RoadPuzzleRailTileData
	SaveDB          RoadPuzzleBoardSaveDB
	ParcelResultDB  ParcelResultDB
}
