package protos

import (
	"time"
)

type ShiftingCraftInfoDB struct {
	SlotSequence int64 `json:",omitempty,omitzero"`
	CraftRecipeId int64 `json:",omitempty,omitzero"`
	CraftAmount int64 `json:",omitempty,omitzero"`
	StartTime time.Time `json:",omitempty,omitzero"`
	EndTime time.Time `json:",omitempty,omitzero"`
}
