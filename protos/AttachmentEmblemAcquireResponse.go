package protos

type AttachmentEmblemAcquireResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EmblemDBs []EmblemDB `json:",omitempty,omitzero"`
}
