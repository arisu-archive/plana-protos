package protos

import (
	"encoding/json"
	"fmt"
)

type SkillLevelBatchGrowthRequestDB struct {
	SkillSlot    SkillLevelBatchGrowthRequestDB_SkillSlotDefault `json:",omitzero"`
	Level        int32                                           `json:",omitempty,omitzero"`
	ReplaceInfos []*SelectTicketReplaceInfo
}

type SkillLevelBatchGrowthRequestDB_SkillSlotDefault SkillSlot

func (s SkillLevelBatchGrowthRequestDB_SkillSlotDefault) IsZero() bool {
	return s == SkillLevelBatchGrowthRequestDB_SkillSlotDefault("None")
}

func (x *SkillLevelBatchGrowthRequestDB) UnmarshalJSON(data []byte) error {
	type aliasSkillLevelBatchGrowthRequestDB SkillLevelBatchGrowthRequestDB
	v := aliasSkillLevelBatchGrowthRequestDB{
		SkillSlot: SkillLevelBatchGrowthRequestDB_SkillSlotDefault("None"),
	}
	if err := json.Unmarshal(data, &v); err != nil { //nolint:musttag // alias inherits json tags from the source struct via `type aliasT T`; linter walks the AST and misses the inheritance.
		return fmt.Errorf("unmarshal SkillLevelBatchGrowthRequestDB: %w", err)
	}
	*x = SkillLevelBatchGrowthRequestDB(v)
	return nil
}
