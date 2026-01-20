package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type MailDB struct {
	ServerId          int64             `json:",omitempty,omitzero"`
	AccountServerId   int64             `json:",omitempty,omitzero"`
	Type              flatdata.MailType `json:",omitempty,omitzero"`
	UniqueId          int64             `json:",omitempty,omitzero"`
	Sender            string            `json:",omitempty,omitzero"`
	Comment           string            `json:",omitempty,omitzero"`
	SendDate          MxTime            `json:",omitempty,omitzero"`
	ReceiptDate       *MxTime           `json:",omitempty,omitzero"`
	ExpireDate        *MxTime           `json:",omitempty,omitzero"`
	OrderId           *int64            `json:",omitempty,omitzero"`
	ProductId         *int64            `json:",omitempty,omitzero"`
	IsCashMail        bool              `json:",omitempty,omitzero"`
	ParcelInfos       []ParcelInfo      `json:",omitempty,omitzero"`
	RemainParcelInfos []ParcelInfo      `json:",omitempty,omitzero"`
}
