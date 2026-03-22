package protos

type ScenarioTacticResultRequest struct {
	RequestPacket
	PassCheckCharacter bool `json:",omitempty,omitzero"`
	Summary            BattleSummary
	Hand               SkillCardHand
	SkipSummary        TacticSkipSummary
}
