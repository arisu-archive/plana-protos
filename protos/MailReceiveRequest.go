package protos

type MailReceiveRequest struct {
	RequestPacket
	MailServerIds []int64 `json:",omitempty,omitzero"`
}
