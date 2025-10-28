package protos

import (
	"time"
)

type CafeDB_CafeCharacterDB struct {
	VisitingCharacterDB
	IsSummon bool `json:",omitempty,omitzero"`
	LastInteractTime time.Time `json:",omitempty,omitzero"`
}
