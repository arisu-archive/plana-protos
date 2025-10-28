package protos

type HexaTileState struct {
	Id int32 `json:",omitempty,omitzero"`
	IsHide bool `json:",omitempty,omitzero"`
	IsFog bool `json:",omitempty,omitzero"`
	CanNotMove bool `json:",omitempty,omitzero"`
}
