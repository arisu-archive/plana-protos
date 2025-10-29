package protos

type MailBoxFullErrorPacket struct {
	ResponsePacket
	ErrorCode WebAPIErrorCode `json:",omitempty,omitzero"`
}
