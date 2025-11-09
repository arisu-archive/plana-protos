package protos

type SkillLevelBatchGrowthRequestDB struct {
	SkillSlot    SkillSlot                 `json:",omitempty,omitzero"`
	Level        int32                     `json:",omitempty,omitzero"`
	ReplaceInfos []SelectTicketReplaceInfo `json:",omitempty,omitzero"`
}
