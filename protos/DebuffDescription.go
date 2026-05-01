package protos

import (
	"encoding/json"
	"fmt"
)

type DebuffDescription struct {
	AccountId             int64                              `json:",omitempty,omitzero"`
	LogicEffectTemplateId string                             `json:",omitempty,omitzero"`
	LogicEffectGroupId    string                             `json:",omitempty,omitzero"`
	LogicEffectLevel      int32                              `json:",omitempty,omitzero"`
	DurationFrame         int32                              `json:",omitempty,omitzero"`
	SkillSlot             DebuffDescription_SkillSlotDefault `json:",omitzero"`
	IssuedTimestamp       int32                              `json:",omitempty,omitzero"`
}

type DebuffDescription_SkillSlotDefault SkillSlot

func (s DebuffDescription_SkillSlotDefault) IsZero() bool {
	return s == DebuffDescription_SkillSlotDefault("None")
}

func (x *DebuffDescription) UnmarshalJSON(data []byte) error {
	type aliasDebuffDescription DebuffDescription
	v := aliasDebuffDescription{
		SkillSlot: DebuffDescription_SkillSlotDefault("None"),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("unmarshal DebuffDescription: %w", err)
	}
	*x = DebuffDescription(v)
	return nil
}
