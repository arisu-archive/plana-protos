package protos

type MiniGameTankRaceSummary struct {
	IsClear            bool   `json:",omitempty,omitzero"`
	TankGroupId        int64  `json:",omitempty,omitzero"`
	TankId             int64  `json:",omitempty,omitzero"`
	Score              int64  `json:",omitempty,omitzero"`
	ClearTimeInSeconds int64  `json:",omitempty,omitzero"`
	MaxCombo           int32  `json:",omitempty,omitzero"`
	BossTankGroupId    *int64 `json:",omitempty,omitzero"`
	DamageTakenCount   *int32 `json:",omitempty,omitzero"`
}
