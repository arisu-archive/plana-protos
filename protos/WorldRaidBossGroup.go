package protos

type WorldRaidBossGroup struct {
	ContentsValueChangeDB
	GroupId       int64  `json:",omitempty,omitzero"`
	BossSpawnTime MxTime `json:",omitempty,omitzero"`
	EliminateTime MxTime `json:",omitempty,omitzero"`
}
