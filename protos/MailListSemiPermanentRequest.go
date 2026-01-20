package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type MailListSemiPermanentRequest struct {
	RequestPacket
	IsReadMail      bool                     `json:",omitempty,omitzero"`
	PivotTime       MxTime                   `json:",omitempty,omitzero"`
	PivotIndex      int64                    `json:",omitempty,omitzero"`
	MailSortingRule flatdata.MailSortingRule `json:",omitempty,omitzero"`
	IsDescending    bool                     `json:",omitempty,omitzero"`
}
