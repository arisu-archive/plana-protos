package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ClanConferRequest struct {
	RequestPacket
	MemberAccountId int64                    `json:",omitempty,omitzero"`
	ConferingGrade  flatdata.ClanSocialGrade `json:",omitempty,omitzero"`
}
