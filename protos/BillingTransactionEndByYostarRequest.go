package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type BillingTransactionEndByYostarRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PurchaseOrderId int64 `json:",omitempty,omitzero"`
	EndType flatdata.BillingTransactionEndType `json:",omitempty,omitzero"`
}
