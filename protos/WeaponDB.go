package protos

import (
	"encoding/json"
	"fmt"
)

type WeaponDB struct {
	ParcelBase
	UniqueId               int64                     `json:",omitempty,omitzero"`
	Level                  WeaponDB_LevelDefault     `json:",omitzero"`
	Exp                    WeaponDB_ExpDefault       `json:",omitzero"`
	StarGrade              WeaponDB_StarGradeDefault `json:",omitzero"`
	BoundCharacterServerId int64                     `json:",omitempty,omitzero"`
}

type WeaponDB_LevelDefault int32

func (s WeaponDB_LevelDefault) IsZero() bool {
	return s == WeaponDB_LevelDefault(1)
}

type WeaponDB_ExpDefault int64

func (s WeaponDB_ExpDefault) IsZero() bool {
	return s == WeaponDB_ExpDefault(0)
}

type WeaponDB_StarGradeDefault int32

func (s WeaponDB_StarGradeDefault) IsZero() bool {
	return s == WeaponDB_StarGradeDefault(1)
}

func (x *WeaponDB) UnmarshalJSON(data []byte) error {
	type aliasWeaponDB WeaponDB
	v := aliasWeaponDB{
		Level:     WeaponDB_LevelDefault(1),
		Exp:       WeaponDB_ExpDefault(0),
		StarGrade: WeaponDB_StarGradeDefault(1),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("unmarshal WeaponDB: %w", err)
	}
	*x = WeaponDB(v)
	return nil
}
