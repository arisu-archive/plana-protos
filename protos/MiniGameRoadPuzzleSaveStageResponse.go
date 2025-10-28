package protos

type MiniGameRoadPuzzleSaveStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDB RoadPuzzleBoardSaveDB `json:",omitempty,omitzero"`
}
