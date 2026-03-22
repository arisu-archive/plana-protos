package protos

type MiniGameCCGSelectRewardCardResponse struct {
	ResponsePacket
	StageDB           MiniGameCCGStagePlayDB
	SaveDB            MiniGameCCGSaveDB
	ReceivedRewardIds []int64
}
