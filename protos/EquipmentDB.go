package protos

import (
	"encoding/json"
	"fmt"
)

type EquipmentDB struct {
	ConsumableItemBaseDB
	Level                  EquipmentDBLevelDefault `json:",omitzero"`
	Exp                    EquipmentDBExpDefault   `json:",omitzero"`
	Tier                   EquipmentDBTierDefault  `json:",omitzero"`
	BoundCharacterServerId int64                   `json:",omitempty,omitzero"`
}

type EquipmentDBLevelDefault int32

func (s EquipmentDBLevelDefault) IsZero() bool {
	return s == EquipmentDBLevelDefault(1)
}

type EquipmentDBExpDefault int64

func (s EquipmentDBExpDefault) IsZero() bool {
	return s == EquipmentDBExpDefault(0)
}

type EquipmentDBTierDefault int32

func (s EquipmentDBTierDefault) IsZero() bool {
	return s == EquipmentDBTierDefault(1)
}

func (x *EquipmentDB) UnmarshalJSON(data []byte) error {
	type aliasEquipmentDB EquipmentDB
	v := aliasEquipmentDB{
		Level: EquipmentDBLevelDefault(1),
		Exp:   EquipmentDBExpDefault(0),
		Tier:  EquipmentDBTierDefault(1),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return fmt.Errorf("unmarshal EquipmentDB: %w", err)
	}
	*x = EquipmentDB(v)
	return nil
}
