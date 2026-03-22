package protos

type CampaignTacticResultRequest struct {
	RequestPacket
	PassCheckCharacter bool `json:",omitempty,omitzero"`
	Summary            BattleSummary
	Hand               SkillCardHand
	SkipSummary        TacticSkipSummary
}
