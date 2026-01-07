package protos

type DebuffDescription struct {
	AccountId             int64     `json:",omitempty,omitzero"`
	LogicEffectTemplateId string    `json:",omitempty,omitzero"`
	LogicEffectGroupId    string    `json:",omitempty,omitzero"`
	LogicEffectLevel      int32     `json:",omitempty,omitzero"`
	DurationFrame         int32     `json:",omitempty,omitzero"`
	SkillSlot             SkillSlot `json:",omitempty,omitzero"`
	IssuedTimestamp       int32     `json:",omitempty,omitzero"`
}
