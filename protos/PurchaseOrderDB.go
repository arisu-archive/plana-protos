package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type PurchaseOrderDB struct {
	ShopCashId int64 `json:",omitempty,omitzero"`
	StatusCode flatdata.PurchaseStatusCode `json:",omitempty,omitzero"`
	PurchaseOrderId int64 `json:",omitempty,omitzero"`
}
