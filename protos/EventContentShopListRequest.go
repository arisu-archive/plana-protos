package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EventContentShopListRequest struct {
	RequestPacket
	EventContentId int64                       `json:",omitempty,omitzero"`
	CategoryList   []flatdata.ShopCategoryType `json:",omitempty,omitzero"`
}
