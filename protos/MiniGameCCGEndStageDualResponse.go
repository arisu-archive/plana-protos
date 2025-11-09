package protos

type MiniGameCCGEndStageDualResponse struct {
	ResponsePacket
	StageDB MiniGameCCGStagePlayDB `json:",omitempty,omitzero"`
	SaveDB  MiniGameCCGSaveDB      `json:",omitempty,omitzero"`
}
