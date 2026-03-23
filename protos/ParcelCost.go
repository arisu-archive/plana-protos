package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ParcelCost struct {
	ParcelInfos      []ParcelInfo
	Currency         *CurrencyTransaction `json:",omitempty,omitzero"`
	EquipmentDBs     []EquipmentDB
	ItemDBs          []ItemDB
	FurnitureDBs     []FurnitureDB
	ConsumeCondition flatdata.ConsumeCondition `json:",omitempty,omitzero"`
}
