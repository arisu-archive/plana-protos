package protos

type AttachmentStudentFrameAttachResponse struct {
	ResponsePacket
	AttachmentDB *AccountAttachmentDB `json:",omitempty,omitzero"`
}
