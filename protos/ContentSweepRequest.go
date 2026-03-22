package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ContentSweepRequest struct {
	RequestPacket
	Content        flatdata.ContentType
	StageId        int64 `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	Count          int64 `json:",omitempty,omitzero"`
}
