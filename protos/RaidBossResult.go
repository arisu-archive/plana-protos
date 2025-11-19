package protos

type RaidBossResult struct {
	RaidDamage         RaidDamage `json:",omitempty,omitzero"`
	EndHpRateRawValue  int64      `json:",omitempty,omitzero"`
	GroggyRateRawValue int64      `json:",omitempty,omitzero"`
	GroggyCount        int32      `json:",omitempty,omitzero"`
	SubPartsHPs        []int64    `json:",omitempty,omitzero"`
	AIPhase            int64      `json:",omitempty,omitzero"`
}
