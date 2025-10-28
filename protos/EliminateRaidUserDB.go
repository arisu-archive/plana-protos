package protos

type EliminateRaidUserDB struct {
	RaidUserDB
	BossGroupToRankingPoint map[string]int64 `json:",omitempty,omitzero"`
}
