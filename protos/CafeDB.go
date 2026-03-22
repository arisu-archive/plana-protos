package protos

import (
	"github.com/arisu-archive/mapx"
)

type CafeDB struct {
	CafeDBId              int64   `json:",omitempty,omitzero"`
	CafeId                int64   `json:",omitempty,omitzero"`
	AccountId             int64   `json:",omitempty,omitzero"`
	CafeRank              int32   `json:",omitempty,omitzero"`
	LastUpdate            MxTime  `json:",omitempty,omitzero"`
	LastSummonDate        *MxTime `json:",omitempty,omitzero"`
	IsNew                 bool    `json:",omitempty,omitzero"`
	CafeVisitCharacterDBs *mapx.OrderedMap[int64, CafeDB_CafeCharacterDB]
	FurnitureDBs          []FurnitureDB
	ProductionAppliedTime MxTime `json:",omitempty,omitzero"`
	ProductionDB          CafeProductionDB
}
