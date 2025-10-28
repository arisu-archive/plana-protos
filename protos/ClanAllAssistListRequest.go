package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ClanAllAssistListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EchelonType flatdata.EchelonType `json:",omitempty,omitzero"`
	PendingAssistUseInfo []ClanAssistUseInfo `json:",omitempty,omitzero"`
	IsPractice bool `json:",omitempty,omitzero"`
}
