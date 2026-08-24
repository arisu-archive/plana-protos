package protos

type FindGiftSummary struct {
	UniqueName string
	ClearCount int32 `json:",omitempty,omitzero"`
}
