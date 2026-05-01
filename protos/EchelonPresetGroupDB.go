package protos

import (
	"encoding/json"
	"fmt"

	"github.com/arisu-archive/mapx"
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EchelonPresetGroupDB struct {
	GroupIndex    EchelonPresetGroupDB_GroupIndexDefault    `json:",omitzero"`
	ExtensionType EchelonPresetGroupDB_ExtensionTypeDefault `json:",omitzero"`
	GroupLabel    string                                    `json:",omitempty,omitzero"`
	PresetDBs     *mapx.OrderedMap[int32, *EchelonPresetDB]
}

type EchelonPresetGroupDB_GroupIndexDefault int32

func (s EchelonPresetGroupDB_GroupIndexDefault) IsZero() bool {
	return s == EchelonPresetGroupDB_GroupIndexDefault(-1)
}

type EchelonPresetGroupDB_ExtensionTypeDefault flatdata.EchelonExtensionType

func (s EchelonPresetGroupDB_ExtensionTypeDefault) IsZero() bool {
	return s == EchelonPresetGroupDB_ExtensionTypeDefault(flatdata.EchelonExtensionTypeBase)
}

func (x *EchelonPresetGroupDB) UnmarshalJSON(data []byte) error {
	type aliasEchelonPresetGroupDB EchelonPresetGroupDB
	v := aliasEchelonPresetGroupDB{
		GroupIndex:    EchelonPresetGroupDB_GroupIndexDefault(-1),
		ExtensionType: EchelonPresetGroupDB_ExtensionTypeDefault(flatdata.EchelonExtensionTypeBase),
	}
	if err := json.Unmarshal(data, &v); err != nil { //nolint:musttag // alias inherits json tags from the source struct via `type aliasT T`; linter walks the AST and misses the inheritance.
		return fmt.Errorf("unmarshal EchelonPresetGroupDB: %w", err)
	}
	*x = EchelonPresetGroupDB(v)
	return nil
}
