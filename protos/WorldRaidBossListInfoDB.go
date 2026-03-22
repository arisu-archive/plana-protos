package protos

type WorldRaidBossListInfoDB struct {
	GroupId      int64 `json:",omitempty,omitzero"`
	WorldBossDB  WorldRaidWorldBossDB
	LocalBossDBs []WorldRaidLocalBossDB
}
