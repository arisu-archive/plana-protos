package protos

type RaidMemberDescription struct {
	AccountId        int64 `json:",omitempty,omitzero"`
	AccountName      string
	CharacterId      int64 `json:",omitempty,omitzero"`
	DamageCollection RaidDamageCollection
}
