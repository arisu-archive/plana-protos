package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EchelonPresetGroupDB struct {
	GroupIndex    int32                         `json:",omitempty,omitzero"`
	ExtensionType flatdata.EchelonExtensionType `json:",omitempty,omitzero"`
	GroupLabel    string                        `json:",omitempty,omitzero"`
	PresetDBs     map[int32]EchelonPresetDB     `json:",omitempty,omitzero"`
}
