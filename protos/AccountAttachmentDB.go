package protos

type AccountAttachmentDB struct {
	AccountId      int64 `json:",omitempty,omitzero"`
	EmblemUniqueId int64 `json:",omitempty,omitzero"`
}
