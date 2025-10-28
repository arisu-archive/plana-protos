package protos

type TimeAttackDungeonCharacterDB struct {
	ServerId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	CostumeId int64 `json:",omitempty,omitzero"`
	StarGrade int32 `json:",omitempty,omitzero"`
	Level int32 `json:",omitempty,omitzero"`
	HasWeapon bool `json:",omitempty,omitzero"`
	WeaponDB WeaponDB `json:",omitempty,omitzero"`
	IsAssist bool `json:",omitempty,omitzero"`
	CombatStyleIndex int32 `json:",omitempty,omitzero"`
}
