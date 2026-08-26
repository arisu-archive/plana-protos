package protos

import (
	"encoding/json"
	"fmt"
)

type WeaponDB struct {
	ParcelBase
	UniqueId               int64                    `json:",omitempty,omitzero"`
	Level                  WeaponDBLevelDefault     `json:",omitzero"`
	Exp                    WeaponDBExpDefault       `json:",omitzero"`
	StarGrade              WeaponDBStarGradeDefault `json:",omitzero"`
	BoundCharacterServerId int64                    `json:",omitempty,omitzero"`
}

type WeaponDBLevelDefault int32

func (s WeaponDBLevelDefault) IsZero() bool {
	return s == WeaponDBLevelDefault(1)
}

type WeaponDBExpDefault int64

func (s WeaponDBExpDefault) IsZero() bool {
	return s == WeaponDBExpDefault(0)
}

type WeaponDBStarGradeDefault int32

func (s WeaponDBStarGradeDefault) IsZero() bool {
	return s == WeaponDBStarGradeDefault(1)
}

func (x *WeaponDB) UnmarshalJSON(data []byte) error {
	type aliasWeaponDB WeaponDB
	v := aliasWeaponDB{
		Level:     WeaponDBLevelDefault(1),
		Exp:       WeaponDBExpDefault(0),
		StarGrade: WeaponDBStarGradeDefault(1),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("unmarshal WeaponDB: %w", err)
	}
	*x = WeaponDB(v)
	return nil
}
