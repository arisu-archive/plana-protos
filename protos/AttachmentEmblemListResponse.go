package protos

type AttachmentEmblemListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EmblemDBs []EmblemDB `json:",omitempty,omitzero"`
}
