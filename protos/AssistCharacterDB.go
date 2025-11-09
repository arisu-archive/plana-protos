package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type AssistCharacterDB struct {
	CharacterDB
	EchelonType             flatdata.EchelonType `json:",omitempty,omitzero"`
	SlotNumber              int32                `json:",omitempty,omitzero"`
	AccountId               int64                `json:",omitempty,omitzero"`
	AssistRelation          AssistRelation       `json:",omitempty,omitzero"`
	AssistCharacterServerId int64                `json:",omitempty,omitzero"`
	NickName                string               `json:",omitempty,omitzero"`
	EquipmentDBs            []EquipmentDB        `json:",omitempty,omitzero"`
	WeaponDB                WeaponDB             `json:",omitempty,omitzero"`
	GearDB                  GearDB               `json:",omitempty,omitzero"`
	CostumeId               int64                `json:",omitempty,omitzero"`
	CostumeDB               CostumeDB            `json:",omitempty,omitzero"`
	IsMulligan              bool                 `json:",omitempty,omitzero"`
	IsTSAInteraction        bool                 `json:",omitempty,omitzero"`
	CombatStyleIndex        int32                `json:",omitempty,omitzero"`
}
