package protos

type CharacterBatchSkillLevelUpdateRequest struct {
	RequestPacket
	TargetCharacterDBId        int64                            `json:",omitempty,omitzero"`
	SkillLevelUpdateRequestDBs []SkillLevelBatchGrowthRequestDB `json:",omitempty,omitzero"`
}
