package protos

type CampaignTutorialStageResultRequest struct {
	RequestPacket
	Summary BattleSummary `json:",omitempty,omitzero"`
}
