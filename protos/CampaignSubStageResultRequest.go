package protos

type CampaignSubStageResultRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PassCheckCharacter bool `json:",omitempty,omitzero"`
	Summary BattleSummary `json:",omitempty,omitzero"`
}
