package protos

type AttachmentEmblemAcquireResponse struct {
	ResponsePacket
	EmblemDBs []EmblemDB `json:",omitempty,omitzero"`
}
