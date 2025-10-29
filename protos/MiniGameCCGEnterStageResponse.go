package protos

type MiniGameCCGEnterStageResponse struct {
	ResponsePacket
	StageDB MiniGameCCGStagePlayDB `json:",omitempty,omitzero"`
}
