package protos

type MailReceiveSemiPermanentRequest struct {
	RequestPacket
	ProductId *int64 `json:",omitempty,omitzero"`
	MailDBId  int64  `json:",omitempty,omitzero"`
}
