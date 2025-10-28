package protos

type CampaignTutorialStageResultRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Summary BattleSummary `json:",omitempty,omitzero"`
}
