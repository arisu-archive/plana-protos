package protos

type EventContentLocationGetInfoResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentLocationDB EventContentLocationDB `json:",omitempty,omitzero"`
}
