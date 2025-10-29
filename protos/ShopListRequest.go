package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ShopListRequest struct {
	RequestPacket
	CategoryList []flatdata.ShopCategoryType `json:",omitempty,omitzero"`
}
