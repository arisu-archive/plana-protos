package protos

type ArenaSummary struct {
	ArenaMapId int64 `json:",omitempty,omitzero"`
	EnemyAccountId int64 `json:",omitempty,omitzero"`
	EnemyAccountLevel int64 `json:",omitempty,omitzero"`
}
