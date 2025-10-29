package protos

type MiniGameRoadPuzzleGetInfoResponse struct {
	ResponsePacket
	SaveDB RoadPuzzleBoardSaveDB `json:",omitempty,omitzero"`
}
