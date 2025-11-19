package protos

type MinigameRhythmSummary struct {
	MusicTitle           string                `json:",omitempty,omitzero"`
	PatternDifficulty    int32                 `json:",omitempty,omitzero"`
	IsSpecial            bool                  `json:",omitempty,omitzero"`
	TotalNoteCount       int32                 `json:",omitempty,omitzero"`
	CriticalCount        int32                 `json:",omitempty,omitzero"`
	AttackCount          int32                 `json:",omitempty,omitzero"`
	MissCount            int32                 `json:",omitempty,omitzero"`
	IsFullCombo          bool                  `json:",omitempty,omitzero"`
	MaxCombo             int32                 `json:",omitempty,omitzero"`
	FinalScore           int64                 `json:",omitempty,omitzero"`
	HPBonusScore         int64                 `json:",omitempty,omitzero"`
	GameStartTime        MxTime                `json:",omitempty,omitzero"`
	GameEndTime          MxTime                `json:",omitempty,omitzero"`
	RhythmGamePlayTime   float32               `json:",omitempty,omitzero"`
	StdDev               float32               `json:",omitempty,omitzero"`
	MinigameJudgeRecords []MinigameJudgeRecord `json:",omitempty,omitzero"`
	IsAutoPlay           bool                  `json:",omitempty,omitzero"`
}
