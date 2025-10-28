package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type MemoryLobbyDB struct {
	ParcelBase
	Type flatdata.ParcelType `json:",omitempty,omitzero"`
	MemoryLobbyUniqueId int64 `json:",omitempty,omitzero"`
}
