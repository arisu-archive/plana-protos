package protos

type ErrorPacket struct {
	ResponsePacket
	Reason    string
	ErrorCode WebAPIErrorCode `json:",omitempty,omitzero"`
}
