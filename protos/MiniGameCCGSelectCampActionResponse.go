package protos

type MiniGameCCGSelectCampActionResponse struct {
	ResponsePacket
	StageDB MiniGameCCGStagePlayDB
	SaveDB  MiniGameCCGSaveDB
}
