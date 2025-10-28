package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
	"time"
)

type MailDB struct {
	ServerId int64 `json:",omitempty,omitzero"`
	AccountServerId int64 `json:",omitempty,omitzero"`
	Type flatdata.MailType `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	Sender string `json:",omitempty,omitzero"`
	Comment string `json:",omitempty,omitzero"`
	SendDate time.Time `json:",omitempty,omitzero"`
	ReceiptDate *time.Time `json:",omitempty,omitzero"`
	ExpireDate *time.Time `json:",omitempty,omitzero"`
	ParcelInfos []ParcelInfo `json:",omitempty,omitzero"`
	RemainParcelInfos []ParcelInfo `json:",omitempty,omitzero"`
}
