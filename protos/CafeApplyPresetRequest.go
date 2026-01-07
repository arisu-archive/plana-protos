package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type CafeApplyPresetRequest struct {
	RequestPacket
	PresetType            flatdata.CafePresetType `json:",omitempty,omitzero"`
	SlotId                int32                   `json:",omitempty,omitzero"`
	CafeDBId              int64                   `json:",omitempty,omitzero"`
	UseOtherCafeFurniture bool                    `json:",omitempty,omitzero"`
}
