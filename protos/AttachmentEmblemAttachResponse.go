package protos

type AttachmentEmblemAttachResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AttachmentDB AccountAttachmentDB `json:",omitempty,omitzero"`
}
