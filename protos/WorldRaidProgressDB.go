package protos

type WorldRaidProgressDB struct {
	SeasonId          int64            `json:",omitempty,omitzero"`
	PhaseId           int64            `json:",omitempty,omitzero"`
	Maps              map[string]int64 `json:",omitempty,omitzero"`
	CarrierSkills     map[string]int32 `json:",omitempty,omitzero"`
	ClearConditionIds []int64          `json:",omitempty,omitzero"`
}
