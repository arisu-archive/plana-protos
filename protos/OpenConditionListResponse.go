package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type OpenConditionListResponse struct {
	ResponsePacket
	ConditionContents []flatdata.OpenConditionContent `json:",omitempty,omitzero"`
}
