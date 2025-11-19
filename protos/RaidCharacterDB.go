package protos

type RaidCharacterDB struct {
	ServerId         int64 `json:",omitempty,omitzero"`
	UniqueId         int64 `json:",omitempty,omitzero"`
	StarGrade        int32 `json:",omitempty,omitzero"`
	Level            int32 `json:",omitempty,omitzero"`
	SlotIndex        int32 `json:",omitempty,omitzero"`
	AccountId        int64 `json:",omitempty,omitzero"`
	IsAssist         bool  `json:",omitempty,omitzero"`
	HasWeapon        bool  `json:",omitempty,omitzero"`
	WeaponStarGrade  int32 `json:",omitempty,omitzero"`
	CostumeId        int64 `json:",omitempty,omitzero"`
	CombatStyleIndex int32 `json:",omitempty,omitzero"`
}
