package protos

type AttachmentEmblemAcquireRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	UniqueIds []int64 `json:",omitempty,omitzero"`
}
