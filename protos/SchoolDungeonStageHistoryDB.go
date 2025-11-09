package protos

type SchoolDungeonStageHistoryDB struct {
	StageUniqueId int64  `json:",omitempty,omitzero"`
	StarFlags     []bool `json:",omitempty,omitzero"`
}
