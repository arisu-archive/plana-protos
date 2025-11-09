package protos

type ArenaCharacterDB struct {
	ServerId               int64           `json:",omitempty,omitzero"`
	UniqueId               int64           `json:",omitempty,omitzero"`
	StarGrade              int32           `json:",omitempty,omitzero"`
	Level                  int32           `json:",omitempty,omitzero"`
	PublicSkillLevel       int32           `json:",omitempty,omitzero"`
	ExSkillLevel           int32           `json:",omitempty,omitzero"`
	PassiveSkillLevel      int32           `json:",omitempty,omitzero"`
	ExtraPassiveSkillLevel int32           `json:",omitempty,omitzero"`
	LeaderSkillLevel       int32           `json:",omitempty,omitzero"`
	SlotIndex              int32           `json:",omitempty,omitzero"`
	EquipmentDBs           []EquipmentDB   `json:",omitempty,omitzero"`
	FavorRankInfo          map[int64]int64 `json:",omitempty,omitzero"`
	PotentialStats         map[int32]int32 `json:",omitempty,omitzero"`
	CombatStyleIndex       int32           `json:",omitempty,omitzero"`
	WeaponDB               WeaponDB        `json:",omitempty,omitzero"`
	GearDB                 GearDB          `json:",omitempty,omitzero"`
	CostumeDB              CostumeDB       `json:",omitempty,omitzero"`
}
