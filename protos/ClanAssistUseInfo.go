package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ClanAssistUseInfo struct {
	CharacterAccountId int64                `json:",omitempty,omitzero"`
	CharacterDBId      int64                `json:",omitempty,omitzero"`
	EchelonType        flatdata.EchelonType `json:",omitempty,omitzero"`
	SlotNumber         int32                `json:",omitempty,omitzero"`
	AssistRelation     AssistRelation       `json:",omitempty,omitzero"`
	EchelonSlotType    int32                `json:",omitempty,omitzero"`
	EchelonSlotIndex   int32                `json:",omitempty,omitzero"`
	CombatStyleIndex   int32                `json:",omitempty,omitzero"`
	IsMulligan         bool                 `json:",omitempty,omitzero"`
	IsTSAInteraction   bool                 `json:",omitempty,omitzero"`
}
