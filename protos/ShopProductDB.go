package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ShopProductDB struct {
	EventContentId int64 `json:",omitempty,omitzero"`
	ShopExcelId int64 `json:",omitempty,omitzero"`
	Category flatdata.ShopCategoryType `json:",omitempty,omitzero"`
	DisplayOrder int64 `json:",omitempty,omitzero"`
	PurchaseCount int64 `json:",omitempty,omitzero"`
	PurchaseCountLimit int64 `json:",omitempty,omitzero"`
	Price int64 `json:",omitempty,omitzero"`
	ProductType ShopProductType `json:",omitempty,omitzero"`
}
