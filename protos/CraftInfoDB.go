package protos

import (
	"time"
)

type CraftInfoDB struct {
	SlotSequence int64 `json:",omitempty,omitzero"`
	StartTime time.Time `json:",omitempty,omitzero"`
	EndTime time.Time `json:",omitempty,omitzero"`
	CraftSlotOpenDate time.Time `json:",omitempty,omitzero"`
	Nodes []CraftNodeDB `json:",omitempty,omitzero"`
}
