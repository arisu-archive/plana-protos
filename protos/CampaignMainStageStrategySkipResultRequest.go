package protos

type CampaignMainStageStrategySkipResultRequest struct {
	RequestPacket
	PassCheckCharacter bool `json:",omitempty,omitzero"`
	Summary BattleSummary `json:",omitempty,omitzero"`
}
