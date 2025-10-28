package protos

type AttachmentEmblemAttachRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
}
