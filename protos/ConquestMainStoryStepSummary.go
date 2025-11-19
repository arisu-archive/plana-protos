package protos

type ConquestMainStoryStepSummary struct {
	ConqueredTileCount int64 `json:",omitempty,omitzero"`
	AllTileCount       int64 `json:",omitempty,omitzero"`
	IsStepOpen         bool  `json:",omitempty,omitzero"`
}
