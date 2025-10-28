package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type IdCardBackgroundDB struct {
	ParcelBase
	Type flatdata.ParcelType `json:",omitempty,omitzero"`
	ServerId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	ParcelInfos []ParcelInfo `json:",omitempty,omitzero"`
}
