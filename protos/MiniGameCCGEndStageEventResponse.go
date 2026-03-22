package protos

type MiniGameCCGEndStageEventResponse struct {
	ResponsePacket
	StageDB MiniGameCCGStagePlayDB
	SaveDB  MiniGameCCGSaveDB
}
