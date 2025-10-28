package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ShopRefreshRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ShopCategoryType flatdata.ShopCategoryType `json:",omitempty,omitzero"`
}
