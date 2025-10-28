package protos

type MiniGameRoadPuzzleClearStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	IsSkip bool `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
