package protos

type MailCheckRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
