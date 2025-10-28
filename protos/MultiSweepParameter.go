package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type MultiSweepParameter struct {
	EventContentId *int64 `json:",omitempty,omitzero"`
	ContentType flatdata.ContentType `json:",omitempty,omitzero"`
	StageId int64 `json:",omitempty,omitzero"`
	SweepCount int32 `json:",omitempty,omitzero"`
}
