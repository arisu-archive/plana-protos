package protos

import (
	"encoding/json"
	"fmt"

	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EchelonPresetDB struct {
	GroupIndex             EchelonPresetDBGroupIndexDefault `json:",omitzero"`
	Index                  EchelonPresetDBIndexDefault      `json:",omitzero"`
	Label                  string                           `json:",omitempty,omitzero"`
	LeaderUniqueId         int64                            `json:",omitempty,omitzero"`
	TSSInteractionUniqueId int64                            `json:",omitempty,omitzero"`
	StrikerUniqueIds       []int64
	SpecialUniqueIds       []int64
	ApcUniqueIds           []int64
	CombatStyleIndex       []int32
	MulliganUniqueIds      []int64
	ExtensionType          flatdata.EchelonExtensionType `json:",omitempty,omitzero"`
}

type EchelonPresetDBGroupIndexDefault int32

func (s EchelonPresetDBGroupIndexDefault) IsZero() bool {
	return s == EchelonPresetDBGroupIndexDefault(-1)
}

type EchelonPresetDBIndexDefault int32

func (s EchelonPresetDBIndexDefault) IsZero() bool {
	return s == EchelonPresetDBIndexDefault(-1)
}

func (x *EchelonPresetDB) UnmarshalJSON(data []byte) error {
	type aliasEchelonPresetDB EchelonPresetDB
	v := aliasEchelonPresetDB{
		GroupIndex: EchelonPresetDBGroupIndexDefault(-1),
		Index:      EchelonPresetDBIndexDefault(-1),
	}
	if err := json.Unmarshal(data, &v); err != nil { //nolint:musttag // alias inherits json tags from the source struct via `type aliasT T`; linter walks the AST and misses the inheritance.
		return fmt.Errorf("unmarshal EchelonPresetDB: %w", err)
	}
	*x = EchelonPresetDB(v)
	return nil
}
