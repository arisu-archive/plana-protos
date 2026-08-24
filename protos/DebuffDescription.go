package protos

type DebuffDescription struct {
	AccountId             int64 `json:",omitempty,omitzero"`
	LogicEffectTemplateId string
	LogicEffectGroupId    string
	LogicEffectLevel      int32     `json:",omitempty,omitzero"`
	DurationFrame         int32     `json:",omitempty,omitzero"`
	SkillSlot             SkillSlot `json:",omitempty,omitzero"`
	IssuedTimestamp       int32     `json:",omitempty,omitzero"`
}
