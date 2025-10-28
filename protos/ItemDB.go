package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ItemDB struct {
	ConsumableItemBaseDB
	Type flatdata.ParcelType `json:",omitempty,omitzero"`
}
