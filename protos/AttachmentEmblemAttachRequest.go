package protos

type AttachmentEmblemAttachRequest struct {
	RequestPacket
	UniqueId int64 `json:",omitempty,omitzero"`
}
