package protos

type CampaignMainStageStrategySkipResultRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PassCheckCharacter bool `json:",omitempty,omitzero"`
	Summary BattleSummary `json:",omitempty,omitzero"`
}
