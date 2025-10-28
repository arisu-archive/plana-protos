package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type CostumeDB struct {
	ParcelBase
	Type flatdata.ParcelType `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	BoundCharacterServerId int64 `json:",omitempty,omitzero"`
}
