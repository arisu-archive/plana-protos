package protos

type MailReceiveRequest struct {
	RequestPacket
	MailServerIds []int64
}
