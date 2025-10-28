package protos

type RaidMemberDescription struct {
	AccountId int64 `json:",omitempty,omitzero"`
	AccountName string `json:",omitempty,omitzero"`
	CharacterId int64 `json:",omitempty,omitzero"`
	DamageCollection RaidDamageCollection `json:",omitempty,omitzero"`
}
