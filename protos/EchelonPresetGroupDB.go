package protos

import (
	"github.com/arisu-archive/mapx"
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EchelonPresetGroupDB struct {
	GroupIndex    int32                         `json:",omitempty,omitzero"`
	ExtensionType flatdata.EchelonExtensionType `json:",omitempty,omitzero"`
	GroupLabel    string
	PresetDBs     *mapx.OrderedMap[int32, *EchelonPresetDB]
}
