package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ParcelBase struct {
	Type flatdata.ParcelType `json:",omitempty,omitzero"`
}
