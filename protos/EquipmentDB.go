package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EquipmentDB struct {
	ConsumableItemBaseDB
	Type flatdata.ParcelType `json:",omitempty,omitzero"`
	Level int32 `json:",omitempty,omitzero"`
	Exp int64 `json:",omitempty,omitzero"`
	Tier int32 `json:",omitempty,omitzero"`
	BoundCharacterServerId int64 `json:",omitempty,omitzero"`
}
