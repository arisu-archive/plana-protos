package protos

type KillLog struct {
	Frame    int32    `json:",omitempty,omitzero"`
	EntityId EntityId `json:",omitempty,omitzero"`
}
