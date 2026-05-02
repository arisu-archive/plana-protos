package protos

import (
	"encoding/json"
	"fmt"
)

type DebuffDescription struct {
	AccountId             int64                             `json:",omitempty,omitzero"`
	LogicEffectTemplateId string                            `json:",omitempty,omitzero"`
	LogicEffectGroupId    string                            `json:",omitempty,omitzero"`
	LogicEffectLevel      int32                             `json:",omitempty,omitzero"`
	DurationFrame         int32                             `json:",omitempty,omitzero"`
	SkillSlot             DebuffDescriptionSkillSlotDefault `json:",omitzero"`
	IssuedTimestamp       int32                             `json:",omitempty,omitzero"`
}

type DebuffDescriptionSkillSlotDefault SkillSlot

func (s DebuffDescriptionSkillSlotDefault) IsZero() bool {
	return s == DebuffDescriptionSkillSlotDefault("None")
}

func (x *DebuffDescription) UnmarshalJSON(data []byte) error {
	type aliasDebuffDescription DebuffDescription
	v := aliasDebuffDescription{
		SkillSlot: DebuffDescriptionSkillSlotDefault("None"),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("unmarshal DebuffDescription: %w", err)
	}
	*x = DebuffDescription(v)
	return nil
}
