package protos

type BattleSummary struct {
	HashKey            int64              `json:",omitempty,omitzero"`
	IsBossBattle       bool               `json:",omitempty,omitzero"`
	BattleType         BattleTypes        `json:",omitempty,omitzero"`
	StageId            int64              `json:",omitempty,omitzero"`
	GroundId           int64              `json:",omitempty,omitzero"`
	Winner             GroupTag           `json:",omitempty,omitzero"`
	EndType            BattleEndType      `json:",omitempty,omitzero"`
	EndFrame           int32              `json:",omitempty,omitzero"`
	Group01Summary     GroupSummary       `json:",omitempty,omitzero"`
	Group02Summary     GroupSummary       `json:",omitempty,omitzero"`
	WeekDungeonSummary WeekDungeonSummary `json:",omitempty,omitzero"`
	RaidSummary        RaidSummary        `json:",omitempty,omitzero"`
	TouchCountSummary  ExcessiveTouch     `json:",omitempty,omitzero"`
	ArenaSummary       ArenaSummary       `json:",omitempty,omitzero"`
	ContinueCount      int32              `json:",omitempty,omitzero"`
	ElapsedRealtime    float32            `json:",omitempty,omitzero"`
	IsAbort            bool               `json:",omitempty,omitzero"`
	IsDefeatBattle     bool               `json:",omitempty,omitzero"`
}
