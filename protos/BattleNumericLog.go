package protos

import (
	"encoding/json"
	"fmt"
)

type BattleNumericLog struct {
	EntityType            BattleNumericLogEntityTypeDefault `json:",omitzero"`
	Category              BattleLogCategory                 `json:",omitempty,omitzero"`
	Source                BattleLogSourceType               `json:",omitempty,omitzero"`
	CalculatedSum         int64                             `json:",omitempty,omitzero"`
	AppliedSum            int64                             `json:",omitempty,omitzero"`
	Count                 int64                             `json:",omitempty,omitzero"`
	CriticalMultiplierMax int64                             `json:",omitempty,omitzero"`
	CriticalCount         int64                             `json:",omitempty,omitzero"`
	CalculatedMin         int64                             `json:",omitempty,omitzero"`
	CalculatedMax         int64                             `json:",omitempty,omitzero"`
	AppliedMin            int64                             `json:",omitempty,omitzero"`
	AppliedMax            int64                             `json:",omitempty,omitzero"`
}

type BattleNumericLogEntityTypeDefault BattleEntityType

func (s BattleNumericLogEntityTypeDefault) IsZero() bool {
	return s == BattleNumericLogEntityTypeDefault("None")
}

func (x *BattleNumericLog) UnmarshalJSON(data []byte) error {
	type aliasBattleNumericLog BattleNumericLog
	v := aliasBattleNumericLog{
		EntityType: BattleNumericLogEntityTypeDefault("None"),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("unmarshal BattleNumericLog: %w", err)
	}
	*x = BattleNumericLog(v)
	return nil
}
