package protos

type ArenaBattleDB struct {
	ArenaBattleServerId int64         `json:",omitempty,omitzero"`
	Season              int64         `json:",omitempty,omitzero"`
	Group               int64         `json:",omitempty,omitzero"`
	BattleStartTime     MxTime        `json:",omitempty,omitzero"`
	BattleEndTime       MxTime        `json:",omitempty,omitzero"`
	Seed                int64         `json:",omitempty,omitzero"`
	AttackingUserDB     ArenaUserDB   `json:",omitempty,omitzero"`
	DefendingUserDB     ArenaUserDB   `json:",omitempty,omitzero"`
	BattleSummary       BattleSummary `json:",omitempty,omitzero"`
}
