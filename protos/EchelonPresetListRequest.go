package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EchelonPresetListRequest struct {
	RequestPacket
	EchelonExtensionType flatdata.EchelonExtensionType `json:",omitempty,omitzero"`
}
