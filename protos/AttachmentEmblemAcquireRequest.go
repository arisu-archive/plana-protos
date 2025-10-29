package protos

type AttachmentEmblemAcquireRequest struct {
	RequestPacket
	UniqueIds []int64 `json:",omitempty,omitzero"`
}
