package protos

type WorldRaidWorldBossDB struct {
	GroupId      int64 `json:",omitempty,omitzero"`
	HP           int64 `json:",omitempty,omitzero"`
	Participants int64 `json:",omitempty,omitzero"`
}
