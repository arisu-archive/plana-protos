package protos

type RaidDamage struct {
	Index int32 `json:",omitempty,omitzero"`
	GivenDamage int64 `json:",omitempty,omitzero"`
	GivenGroggyPoint int64 `json:",omitempty,omitzero"`
}
