package protos

type AttachmentStudentFrameListResponse struct {
	ResponsePacket
	StudentFrameDBs []*StudentFrameDB
}
