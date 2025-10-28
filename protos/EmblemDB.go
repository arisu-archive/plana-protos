package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
	"time"
)

type EmblemDB struct {
	ParcelBase
	Type flatdata.ParcelType `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	ReceiveDate time.Time `json:",omitempty,omitzero"`
	ParcelInfos []ParcelInfo `json:",omitempty,omitzero"`
}
