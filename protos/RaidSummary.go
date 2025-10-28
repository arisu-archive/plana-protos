package protos

type RaidSummary struct {
	RaidSeasonId int64 `json:",omitempty,omitzero"`
	RaidBossResults RaidBossResultCollection `json:",omitempty,omitzero"`
}
