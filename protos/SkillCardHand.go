package protos

type SkillCardHand struct {
	Cost float32 `json:",omitempty,omitzero"`
	SkillCardsInHand []SkillCardInfo `json:",omitempty,omitzero"`
}
