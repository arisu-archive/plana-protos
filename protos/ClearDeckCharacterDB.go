package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ClearDeckCharacterDB struct {
	UniqueId         int64              `json:",omitempty,omitzero"`
	StarGrade        int32              `json:",omitempty,omitzero"`
	Level            int32              `json:",omitempty,omitzero"`
	SlotNumber       int32              `json:",omitempty,omitzero"`
	HasWeapon        bool               `json:",omitempty,omitzero"`
	SquadType        flatdata.SquadType `json:",omitempty,omitzero"`
	WeaponStarGrade  int32              `json:",omitempty,omitzero"`
	CombatStyleIndex int32              `json:",omitempty,omitzero"`
}
