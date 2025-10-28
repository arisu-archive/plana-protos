package protos

type AttachmentEmblemListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
