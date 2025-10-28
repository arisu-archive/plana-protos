package protos

type WorldRaidLocalBossDB struct {
	SeasonId int64 `json:",omitempty,omitzero"`
	GroupId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	IsScenario bool `json:",omitempty,omitzero"`
	IsCleardEver bool `json:",omitempty,omitzero"`
	TacticMscSum int64 `json:",omitempty,omitzero"`
	RaidBattleDB RaidBattleDB `json:",omitempty,omitzero"`
	IsContinue bool `json:",omitempty,omitzero"`
}
