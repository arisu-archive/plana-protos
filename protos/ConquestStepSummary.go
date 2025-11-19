package protos

type ConquestStepSummary struct {
	ConqueredTileCount    int64 `json:",omitempty,omitzero"`
	AllTileCount          int64 `json:",omitempty,omitzero"`
	ErosionRemainingCount int64 `json:",omitempty,omitzero"`
	HasPhaseComplete      bool  `json:",omitempty,omitzero"`
	IsErosionPhaseStart   bool  `json:",omitempty,omitzero"`
	IsStepOpen            bool  `json:",omitempty,omitzero"`
}
