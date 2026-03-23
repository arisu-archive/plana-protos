package protos

type AttachmentEmblemListResponse struct {
	ResponsePacket
	EmblemDBs []*EmblemDB
}
