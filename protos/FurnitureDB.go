package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type FurnitureDB struct {
	ConsumableItemBaseDB
	Location           flatdata.FurnitureLocation `json:",omitempty,omitzero"`
	CafeDBId           int64                      `json:",omitempty,omitzero"`
	PositionX          float32                    `json:",omitempty,omitzero"`
	PositionY          float32                    `json:",omitempty,omitzero"`
	Rotation           float32                    `json:",omitempty,omitzero"`
	ItemDeploySequence int64                      `json:",omitempty,omitzero"`
}
