package protos

import (
	"github.com/arisu-archive/mapx"
)

type ArenaCharacterDB struct {
	ServerId               int64 `json:",omitempty,omitzero"`
	UniqueId               int64 `json:",omitempty,omitzero"`
	StarGrade              int32 `json:",omitempty,omitzero"`
	Level                  int32 `json:",omitempty,omitzero"`
	PublicSkillLevel       int32 `json:",omitempty,omitzero"`
	ExSkillLevel           int32 `json:",omitempty,omitzero"`
	PassiveSkillLevel      int32 `json:",omitempty,omitzero"`
	ExtraPassiveSkillLevel int32 `json:",omitempty,omitzero"`
	LeaderSkillLevel       int32 `json:",omitempty,omitzero"`
	SlotIndex              int32 `json:",omitempty,omitzero"`
	EquipmentDBs           []EquipmentDB
	FavorRankInfo          *mapx.OrderedMap[int64, int64]
	PotentialStats         *mapx.OrderedMap[int32, int32]
	CombatStyleIndex       int32 `json:",omitempty,omitzero"`
	WeaponDB               WeaponDB
	GearDB                 GearDB
	CostumeDB              CostumeDB
}
