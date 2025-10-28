package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ContentSaveDiscardRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ContentType flatdata.ContentType `json:",omitempty,omitzero"`
	StageUniqueId int64 `json:",omitempty,omitzero"`
}
