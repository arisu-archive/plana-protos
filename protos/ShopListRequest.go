package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ShopListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CategoryList []flatdata.ShopCategoryType `json:",omitempty,omitzero"`
}
