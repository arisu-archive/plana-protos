package protos

type AttachmentEmblemAttachResponse struct {
	ResponsePacket
	AttachmentDB AccountAttachmentDB `json:",omitempty,omitzero"`
}
