package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
	"time"
)

type ClanAssistSlotDB struct {
	EchelonType flatdata.EchelonType `json:",omitempty,omitzero"`
	SlotNumber int64 `json:",omitempty,omitzero"`
	CharacterDBId int64 `json:",omitempty,omitzero"`
	DeployDate time.Time `json:",omitempty,omitzero"`
	TotalRentCount int64 `json:",omitempty,omitzero"`
	CombatStyleIndex int32 `json:",omitempty,omitzero"`
}
