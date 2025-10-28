package protos

import (
	"time"
)

type CafeProductionDB struct {
	CafeDBId int64 `json:",omitempty,omitzero"`
	ComfortValue int64 `json:",omitempty,omitzero"`
	AppliedDate time.Time `json:",omitempty,omitzero"`
	ProductionParcelInfos []CafeProductionDB_CafeProductionParcelInfo `json:",omitempty,omitzero"`
}
