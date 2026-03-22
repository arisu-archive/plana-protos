package protos

type BattleSummary struct {
	HashKey            int64         `json:",omitempty,omitzero"`
	IsBossBattle       bool          `json:",omitempty,omitzero"`
	BattleType         BattleTypes   `json:",omitempty,omitzero"`
	StageId            int64         `json:",omitempty,omitzero"`
	GroundId           int64         `json:",omitempty,omitzero"`
	Winner             string        `json:",omitempty,omitzero"`
	EndType            BattleEndType `json:",omitempty,omitzero"`
	EndFrame           int32         `json:",omitempty,omitzero"`
	Group01Summary     GroupSummary
	Group02Summary     GroupSummary
	WeekDungeonSummary WeekDungeonSummary
	RaidSummary        RaidSummary
	TouchCountSummary  ExcessiveTouch
	ArenaSummary       ArenaSummary
	ContinueCount      int32   `json:",omitempty,omitzero"`
	ElapsedRealtime    float32 `json:",omitempty,omitzero"`
	IsAbort            bool    `json:",omitempty,omitzero"`
	IsDefeatBattle     bool    `json:",omitempty,omitzero"`
}
