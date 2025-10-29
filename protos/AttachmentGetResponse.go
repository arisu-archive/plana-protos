package protos

type AttachmentGetResponse struct {
	ResponsePacket
	AccountAttachmentDB AccountAttachmentDB `json:",omitempty,omitzero"`
}
