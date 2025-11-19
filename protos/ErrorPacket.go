package protos

type ErrorPacket struct {
	ResponsePacket
	Reason    string          `json:",omitempty,omitzero"`
	ErrorCode WebAPIErrorCode `json:",omitempty,omitzero"`
}
