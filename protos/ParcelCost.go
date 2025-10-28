package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ParcelCost struct {
	ParcelInfos []ParcelInfo `json:",omitempty,omitzero"`
	Currency CurrencyTransaction `json:",omitempty,omitzero"`
	EquipmentDBs []EquipmentDB `json:",omitempty,omitzero"`
	ItemDBs []ItemDB `json:",omitempty,omitzero"`
	FurnitureDBs []FurnitureDB `json:",omitempty,omitzero"`
	ConsumeCondition flatdata.ConsumeCondition `json:",omitempty,omitzero"`
}
