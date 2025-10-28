package protos

type MailReceiveRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MailServerIds []int64 `json:",omitempty,omitzero"`
}
