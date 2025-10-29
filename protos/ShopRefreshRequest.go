package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ShopRefreshRequest struct {
	RequestPacket
	ShopCategoryType flatdata.ShopCategoryType `json:",omitempty,omitzero"`
}
