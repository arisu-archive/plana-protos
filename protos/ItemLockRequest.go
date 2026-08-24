package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ItemLockRequest struct {
	RequestPacket
	TargetType flatdata.ParcelType `json:",omitempty,omitzero"`
	UniqueIds  []int64
}
