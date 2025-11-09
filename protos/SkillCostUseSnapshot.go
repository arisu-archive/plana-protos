package protos

type SkillCostUseSnapshot struct {
	Frame  int64   `json:",omitempty,omitzero"`
	Used   float32 `json:",omitempty,omitzero"`
	CharId int64   `json:",omitempty,omitzero"`
	Level  int32   `json:",omitempty,omitzero"`
}
