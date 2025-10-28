package protos

import (
	"time"
)

type CafeDB struct {
	CafeDBId int64 `json:",omitempty,omitzero"`
	CafeId int64 `json:",omitempty,omitzero"`
	AccountId int64 `json:",omitempty,omitzero"`
	CafeRank int32 `json:",omitempty,omitzero"`
	LastUpdate time.Time `json:",omitempty,omitzero"`
	LastSummonDate *time.Time `json:",omitempty,omitzero"`
	IsNew bool `json:",omitempty,omitzero"`
	CafeVisitCharacterDBs map[int64]CafeDB_CafeCharacterDB `json:",omitempty,omitzero"`
	FurnitureDBs []FurnitureDB `json:",omitempty,omitzero"`
	ProductionAppliedTime time.Time `json:",omitempty,omitzero"`
	ProductionDB CafeProductionDB `json:",omitempty,omitzero"`
}
