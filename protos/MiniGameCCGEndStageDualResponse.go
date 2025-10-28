package protos

type MiniGameCCGEndStageDualResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	StageDB MiniGameCCGStagePlayDB `json:",omitempty,omitzero"`
	SaveDB MiniGameCCGSaveDB `json:",omitempty,omitzero"`
}
