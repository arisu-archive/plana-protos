package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
	"time"
)

type MailListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	IsReadMail bool `json:",omitempty,omitzero"`
	PivotTime time.Time `json:",omitempty,omitzero"`
	PivotIndex int64 `json:",omitempty,omitzero"`
	MailSortingRule flatdata.MailSortingRule `json:",omitempty,omitzero"`
	IsDescending bool `json:",omitempty,omitzero"`
}
