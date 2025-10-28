package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type SelectTicketReplaceInfo struct {
	MaterialType flatdata.ParcelType `json:",omitempty,omitzero"`
	MaterialId int64 `json:",omitempty,omitzero"`
	TicketItemId int64 `json:",omitempty,omitzero"`
	Amount int32 `json:",omitempty,omitzero"`
}
