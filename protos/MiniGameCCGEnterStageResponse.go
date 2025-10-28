package protos

type MiniGameCCGEnterStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StageDB MiniGameCCGStagePlayDB `json:",omitempty,omitzero"`
}
