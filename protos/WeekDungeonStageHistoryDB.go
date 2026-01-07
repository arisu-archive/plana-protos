package protos

type WeekDungeonStageHistoryDB struct {
	AccountServerId int64            `json:",omitempty,omitzero"`
	StageUniqueId   int64            `json:",omitempty,omitzero"`
	StarGoalRecord  map[string]int64 `json:",omitempty,omitzero"`
}
