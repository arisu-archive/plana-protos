package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EventRewardIncreaseDB struct {
	EventTargetType flatdata.EventTargetType `json:",omitempty,omitzero"`
	Multiplier      *BasisPoint              `json:",omitempty,omitzero"`
	BeginDate       MxTime                   `json:",omitempty,omitzero"`
	EndDate         MxTime                   `json:",omitempty,omitzero"`
}
