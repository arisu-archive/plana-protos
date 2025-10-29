package protos

type EventContentLocationGetInfoResponse struct {
	ResponsePacket
	EventContentLocationDB EventContentLocationDB `json:",omitempty,omitzero"`
}
