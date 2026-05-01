package protos

import (
	"encoding/json"
	"fmt"

	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EchelonPresetDB struct {
	GroupIndex             EchelonPresetDB_GroupIndexDefault `json:",omitzero"`
	Index                  EchelonPresetDB_IndexDefault      `json:",omitzero"`
	Label                  string                            `json:",omitempty,omitzero"`
	LeaderUniqueId         int64                             `json:",omitempty,omitzero"`
	TSSInteractionUniqueId int64                             `json:",omitempty,omitzero"`
	StrikerUniqueIds       []int64
	SpecialUniqueIds       []int64
	CombatStyleIndex       []int32
	MulliganUniqueIds      []int64
	ExtensionType          flatdata.EchelonExtensionType `json:",omitempty,omitzero"`
}

type EchelonPresetDB_GroupIndexDefault int32

func (s EchelonPresetDB_GroupIndexDefault) IsZero() bool {
	return s == EchelonPresetDB_GroupIndexDefault(-1)
}

type EchelonPresetDB_IndexDefault int32

func (s EchelonPresetDB_IndexDefault) IsZero() bool {
	return s == EchelonPresetDB_IndexDefault(-1)
}

func (x *EchelonPresetDB) UnmarshalJSON(data []byte) error {
	type aliasEchelonPresetDB EchelonPresetDB
	v := aliasEchelonPresetDB{
		GroupIndex: EchelonPresetDB_GroupIndexDefault(-1),
		Index:      EchelonPresetDB_IndexDefault(-1),
	}
	if err := json.Unmarshal(data, &v); err != nil { //nolint:musttag // alias inherits json tags from the source struct via `type aliasT T`; linter walks the AST and misses the inheritance.
		return fmt.Errorf("unmarshal EchelonPresetDB: %w", err)
	}
	*x = EchelonPresetDB(v)
	return nil
}
