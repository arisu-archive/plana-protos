package protos

type EventImageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ImageBytes []byte `json:",omitempty,omitzero"`
}
