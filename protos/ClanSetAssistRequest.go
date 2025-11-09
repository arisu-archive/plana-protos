package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ClanSetAssistRequest struct {
	RequestPacket
	EchelonType      flatdata.EchelonType `json:",omitempty,omitzero"`
	SlotNumber       int32                `json:",omitempty,omitzero"`
	CharacterDBId    int64                `json:",omitempty,omitzero"`
	CombatStyleIndex int32                `json:",omitempty,omitzero"`
}
