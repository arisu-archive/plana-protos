package protos

type Strategy struct {
	EntityId int64 `json:",omitempty,omitzero"`
	Rotate Vector3 `json:",omitempty,omitzero"`
	Id int64 `json:",omitempty,omitzero"`
	Location HexLocation `json:",omitempty,omitzero"`
	PlayAnimation bool `json:",omitempty,omitzero"`
	Activated bool `json:",omitempty,omitzero"`
	Values []int32 `json:",omitempty,omitzero"`
	Index int32 `json:",omitempty,omitzero"`
}
