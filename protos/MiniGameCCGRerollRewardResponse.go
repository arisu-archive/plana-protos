package protos

type MiniGameCCGRerollRewardResponse struct {
	ResponsePacket
	StageDB MiniGameCCGStagePlayDB `json:",omitempty,omitzero"`
}
