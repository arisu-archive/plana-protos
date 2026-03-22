package protos

type EventContentTacticResultRequest struct {
	RequestPacket
	EventContentId     int64 `json:",omitempty,omitzero"`
	PassCheckCharacter bool  `json:",omitempty,omitzero"`
	Summary            BattleSummary
	Hand               SkillCardHand
	SkipSummary        TacticSkipSummary
}
