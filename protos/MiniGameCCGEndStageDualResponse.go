package protos

type MiniGameCCGEndStageDualResponse struct {
	ResponsePacket
	StageDB MiniGameCCGStagePlayDB
	SaveDB  MiniGameCCGSaveDB
}
