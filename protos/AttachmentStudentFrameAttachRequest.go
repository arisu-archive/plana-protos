package protos

type AttachmentStudentFrameAttachRequest struct {
	RequestPacket
	UniqueId int64 `json:",omitempty,omitzero"`
}
