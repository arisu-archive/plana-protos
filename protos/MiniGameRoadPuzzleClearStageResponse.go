package protos

type MiniGameRoadPuzzleClearStageResponse struct {
	ResponsePacket
	IsSkip         bool `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB
}
