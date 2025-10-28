package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
	"time"
)

type ShopInfoDB struct {
	EventContentId int64 `json:",omitempty,omitzero"`
	Category flatdata.ShopCategoryType `json:",omitempty,omitzero"`
	ManualRefreshCount *int64 `json:",omitempty,omitzero"`
	IsRefresh bool `json:",omitempty,omitzero"`
	ShopGroupType flatdata.ShopGroupType `json:",omitempty,omitzero"`
	NextAutoRefreshDate *time.Time `json:",omitempty,omitzero"`
	LastAutoRefreshDate *time.Time `json:",omitempty,omitzero"`
	ShopProductList []ShopProductDB `json:",omitempty,omitzero"`
}
