package protos

type MiniGameRoadPuzzleSaveStageResponse struct {
	ResponsePacket
	SaveDB RoadPuzzleBoardSaveDB `json:",omitempty,omitzero"`
}
