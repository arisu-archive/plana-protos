package protos

type EquipmentDB struct {
	ConsumableItemBaseDB
	Level                  int32 `json:",omitempty,omitzero"`
	Exp                    int64 `json:",omitempty,omitzero"`
	Tier                   int32 `json:",omitempty,omitzero"`
	BoundCharacterServerId int64 `json:",omitempty,omitzero"`
}
