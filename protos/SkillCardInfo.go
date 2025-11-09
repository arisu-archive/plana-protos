package protos

type SkillCardInfo struct {
	CharacterId    int64  `json:",omitempty,omitzero"`
	HandIndex      int32  `json:",omitempty,omitzero"`
	SkillId        string `json:",omitempty,omitzero"`
	RemainCoolTime int32  `json:",omitempty,omitzero"`
}
