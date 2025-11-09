package protos

type MiniGameCCGCardDB struct {
	CardDBId int32 `json:",omitempty,omitzero"`
	CardId   int64 `json:",omitempty,omitzero"`
}
