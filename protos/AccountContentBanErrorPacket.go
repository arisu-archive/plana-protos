package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type AccountContentBanErrorPacket struct {
	ResponsePacket
	ErrorCode    WebAPIErrorCode      `json:",omitempty,omitzero"`
	ContentType  flatdata.ContentType `json:",omitempty,omitzero"`
	BanStartDate MxTime               `json:",omitempty,omitzero"`
	BanEndDate   MxTime               `json:",omitempty,omitzero"`
}
