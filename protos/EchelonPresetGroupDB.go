package protos

import (
	"encoding/json"
	"fmt"

	"github.com/arisu-archive/mapx"
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EchelonPresetGroupDB struct {
	GroupIndex    EchelonPresetGroupDBGroupIndexDefault    `json:",omitzero"`
	ExtensionType EchelonPresetGroupDBExtensionTypeDefault `json:",omitzero"`
	GroupLabel    string                                   `json:",omitempty,omitzero"`
	PresetDBs     *mapx.OrderedMap[int32, *EchelonPresetDB]
}

type EchelonPresetGroupDBGroupIndexDefault int32

func (s EchelonPresetGroupDBGroupIndexDefault) IsZero() bool {
	return s == EchelonPresetGroupDBGroupIndexDefault(-1)
}

type EchelonPresetGroupDBExtensionTypeDefault flatdata.EchelonExtensionType

func (s EchelonPresetGroupDBExtensionTypeDefault) IsZero() bool {
	return s == EchelonPresetGroupDBExtensionTypeDefault(flatdata.EchelonExtensionTypeBase)
}

func (x *EchelonPresetGroupDB) UnmarshalJSON(data []byte) error {
	type aliasEchelonPresetGroupDB EchelonPresetGroupDB
	v := aliasEchelonPresetGroupDB{
		GroupIndex:    EchelonPresetGroupDBGroupIndexDefault(-1),
		ExtensionType: EchelonPresetGroupDBExtensionTypeDefault(flatdata.EchelonExtensionTypeBase),
	}
	if err := json.Unmarshal(data, &v); err != nil { //nolint:musttag // alias inherits json tags from the source struct via `type aliasT T`; linter walks the AST and misses the inheritance.
		return fmt.Errorf("unmarshal EchelonPresetGroupDB: %w", err)
	}
	*x = EchelonPresetGroupDB(v)
	return nil
}
