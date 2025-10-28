package protos

type SingleRaidUserDB struct {
	RaidUserDB
	RaidTeamSettingDB RaidTeamSettingDB `json:",omitempty,omitzero"`
}
