package protos

type BattleNumericLog struct {
	EntityType            BattleEntityType    `json:",omitempty,omitzero"`
	Category              BattleLogCategory   `json:",omitempty,omitzero"`
	Source                BattleLogSourceType `json:",omitempty,omitzero"`
	CalculatedSum         int64               `json:",omitempty,omitzero"`
	AppliedSum            int64               `json:",omitempty,omitzero"`
	Count                 int64               `json:",omitempty,omitzero"`
	CriticalMultiplierMax int64               `json:",omitempty,omitzero"`
	CriticalCount         int64               `json:",omitempty,omitzero"`
	CalculatedMin         int64               `json:",omitempty,omitzero"`
	CalculatedMax         int64               `json:",omitempty,omitzero"`
	AppliedMin            int64               `json:",omitempty,omitzero"`
	AppliedMax            int64               `json:",omitempty,omitzero"`
}
