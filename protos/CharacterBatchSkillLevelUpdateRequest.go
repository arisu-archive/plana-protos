package protos

type CharacterBatchSkillLevelUpdateRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TargetCharacterDBId int64 `json:",omitempty,omitzero"`
	SkillLevelUpdateRequestDBs []SkillLevelBatchGrowthRequestDB `json:",omitempty,omitzero"`
}
