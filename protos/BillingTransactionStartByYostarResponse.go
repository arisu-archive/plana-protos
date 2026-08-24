package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type BillingTransactionStartByYostarResponse struct {
	ResponsePacket
	PurchaseCount             int64  `json:",omitempty,omitzero"`
	PurchaseResetDate         MxTime `json:",omitempty,omitzero"`
	PurchaseOrderId           int64  `json:",omitempty,omitzero"`
	MXSeedKey                 string
	PurchaseServerTag         flatdata.PurchaseServerTag `json:",omitempty,omitzero"`
	PurchaseServerCallbackUrl string
}
