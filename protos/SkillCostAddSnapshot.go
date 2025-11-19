package protos

type SkillCostAddSnapshot struct {
	Frame int64   `json:",omitempty,omitzero"`
	Added float32 `json:",omitempty,omitzero"`
}
