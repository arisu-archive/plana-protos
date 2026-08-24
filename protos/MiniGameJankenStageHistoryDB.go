package protos

type MiniGameJankenStageHistoryDB struct {
	StageId   int64 `json:",omitempty,omitzero"`
	Star1Flag bool  `json:",omitempty,omitzero"`
	Star2Flag bool  `json:",omitempty,omitzero"`
	Star3Flag bool  `json:",omitempty,omitzero"`
}
