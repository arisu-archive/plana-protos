package protos

import (
	"encoding/json"
	"fmt"
)

type EquipmentDB struct {
	ConsumableItemBaseDB
	Level                  EquipmentDB_LevelDefault `json:",omitzero"`
	Exp                    EquipmentDB_ExpDefault   `json:",omitzero"`
	Tier                   EquipmentDB_TierDefault  `json:",omitzero"`
	BoundCharacterServerId int64                    `json:",omitempty,omitzero"`
}

type EquipmentDB_LevelDefault int32

func (s EquipmentDB_LevelDefault) IsZero() bool {
	return s == EquipmentDB_LevelDefault(1)
}

type EquipmentDB_ExpDefault int64

func (s EquipmentDB_ExpDefault) IsZero() bool {
	return s == EquipmentDB_ExpDefault(0)
}

type EquipmentDB_TierDefault int32

func (s EquipmentDB_TierDefault) IsZero() bool {
	return s == EquipmentDB_TierDefault(1)
}

func (x *EquipmentDB) UnmarshalJSON(data []byte) error {
	type aliasEquipmentDB EquipmentDB
	v := aliasEquipmentDB{
		Level: EquipmentDB_LevelDefault(1),
		Exp:   EquipmentDB_ExpDefault(0),
		Tier:  EquipmentDB_TierDefault(1),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("unmarshal EquipmentDB: %w", err)
	}
	*x = EquipmentDB(v)
	return nil
}
