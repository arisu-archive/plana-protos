package protos

type AttachmentGetResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	AccountAttachmentDB AccountAttachmentDB `json:",omitempty,omitzero"`
}
