package protos

type FindGiftSummary struct {
	UniqueName string `json:",omitempty,omitzero"`
	ClearCount int32  `json:",omitempty,omitzero"`
}
