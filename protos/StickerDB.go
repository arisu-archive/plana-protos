package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type StickerDB struct {
	ParcelBase
	Type flatdata.ParcelType `json:",omitempty,omitzero"`
	StickerUniqueId int64 `json:",omitempty,omitzero"`
	ParcelInfos []ParcelInfo `json:",omitempty,omitzero"`
}
