package protos

type SkipHistoryDB struct {
	Prologue int32 `json:",omitempty,omitzero"`
	Tutorial map[int32]int32 `json:",omitempty,omitzero"`
}
