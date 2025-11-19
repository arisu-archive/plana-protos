package protos

type MiniGameCCGSelectRewardCardResponse struct {
	ResponsePacket
	StageDB           MiniGameCCGStagePlayDB `json:",omitempty,omitzero"`
	SaveDB            MiniGameCCGSaveDB      `json:",omitempty,omitzero"`
	ReceivedRewardIds []int64                `json:",omitempty,omitzero"`
}
