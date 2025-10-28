package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EchelonPresetGroupRenameRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PresetGroupIndex int32 `json:",omitempty,omitzero"`
	ExtensionType flatdata.EchelonExtensionType `json:",omitempty,omitzero"`
	PresetGroupLabel string `json:",omitempty,omitzero"`
}
