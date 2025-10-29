package protos

type EventImageResponse struct {
	ResponsePacket
	ImageBytes []byte `json:",omitempty,omitzero"`
}
