package protos

type MiniGameRoadPuzzleGetInfoResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDB RoadPuzzleBoardSaveDB `json:",omitempty,omitzero"`
}
