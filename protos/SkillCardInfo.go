package protos

type SkillCardInfo struct {
	CharacterId    int64 `json:",omitempty,omitzero"`
	HandIndex      int32 `json:",omitempty,omitzero"`
	SkillId        string
	RemainCoolTime int32 `json:",omitempty,omitzero"`
}
