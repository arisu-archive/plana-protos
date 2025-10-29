package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EventContentShopRefreshRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	ShopCategoryType flatdata.ShopCategoryType `json:",omitempty,omitzero"`
}
