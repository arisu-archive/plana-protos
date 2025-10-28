package protos

type ErrorPacket struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Reason string `json:",omitempty,omitzero"`
	ErrorCode WebAPIErrorCode `json:",omitempty,omitzero"`
}
