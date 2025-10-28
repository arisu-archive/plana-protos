package protos

type AttachmentGetRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
