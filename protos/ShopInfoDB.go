package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ShopInfoDB struct {
	EventContentId      int64                     `json:",omitempty,omitzero"`
	Category            flatdata.ShopCategoryType `json:",omitempty,omitzero"`
	ManualRefreshCount  *int64                    `json:",omitempty,omitzero"`
	IsRefresh           bool                      `json:",omitempty,omitzero"`
	ShopGroupType       flatdata.ShopGroupType    `json:",omitempty,omitzero"`
	NextAutoRefreshDate *MxTime                   `json:",omitempty,omitzero"`
	LastAutoRefreshDate *MxTime                   `json:",omitempty,omitzero"`
	ShopProductList     []ShopProductDB
}
