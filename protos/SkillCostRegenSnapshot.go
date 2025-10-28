package protos

type SkillCostRegenSnapshot struct {
	Frame int64 `json:",omitempty,omitzero"`
	Regen float32 `json:",omitempty,omitzero"`
}
