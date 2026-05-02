package protos

import (
	"encoding/json"
	"fmt"
)

type SkillLevelBatchGrowthRequestDB struct {
	SkillSlot    SkillLevelBatchGrowthRequestDBSkillSlotDefault `json:",omitzero"`
	Level        int32                                          `json:",omitempty,omitzero"`
	ReplaceInfos []*SelectTicketReplaceInfo
}

type SkillLevelBatchGrowthRequestDBSkillSlotDefault SkillSlot

func (s SkillLevelBatchGrowthRequestDBSkillSlotDefault) IsZero() bool {
	return s == SkillLevelBatchGrowthRequestDBSkillSlotDefault("None")
}

func (x *SkillLevelBatchGrowthRequestDB) UnmarshalJSON(data []byte) error {
	type aliasSkillLevelBatchGrowthRequestDB SkillLevelBatchGrowthRequestDB
	v := aliasSkillLevelBatchGrowthRequestDB{
		SkillSlot: SkillLevelBatchGrowthRequestDBSkillSlotDefault("None"),
	}
	if err := json.Unmarshal(data, &v); err != nil { //nolint:musttag // alias inherits json tags from the source struct via `type aliasT T`; linter walks the AST and misses the inheritance.
		return fmt.Errorf("unmarshal SkillLevelBatchGrowthRequestDB: %w", err)
	}
	*x = SkillLevelBatchGrowthRequestDB(v)
	return nil
}
