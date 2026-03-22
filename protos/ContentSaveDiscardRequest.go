package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ContentSaveDiscardRequest struct {
	RequestPacket
	ContentType   flatdata.ContentType
	StageUniqueId int64 `json:",omitempty,omitzero"`
}
