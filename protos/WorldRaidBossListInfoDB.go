package protos

type WorldRaidBossListInfoDB struct {
	GroupId      int64                  `json:",omitempty,omitzero"`
	WorldBossDB  WorldRaidWorldBossDB   `json:",omitempty,omitzero"`
	LocalBossDBs []WorldRaidLocalBossDB `json:",omitempty,omitzero"`
}
