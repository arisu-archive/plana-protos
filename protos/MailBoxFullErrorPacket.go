package protos

type MailBoxFullErrorPacket struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ErrorCode WebAPIErrorCode `json:",omitempty,omitzero"`
}
