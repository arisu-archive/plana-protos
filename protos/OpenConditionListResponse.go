package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type OpenConditionListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ConditionContents []flatdata.OpenConditionContent `json:",omitempty,omitzero"`
}
