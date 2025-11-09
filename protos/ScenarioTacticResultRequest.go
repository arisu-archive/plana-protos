package protos

type ScenarioTacticResultRequest struct {
	RequestPacket
	PassCheckCharacter bool              `json:",omitempty,omitzero"`
	Summary            BattleSummary     `json:",omitempty,omitzero"`
	Hand               SkillCardHand     `json:",omitempty,omitzero"`
	SkipSummary        TacticSkipSummary `json:",omitempty,omitzero"`
}
