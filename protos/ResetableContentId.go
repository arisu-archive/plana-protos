package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ResetableContentId struct {
	Type flatdata.ResetContentType `json:",omitempty,omitzero"`
	Mapped int64 `json:",omitempty,omitzero"`
}
