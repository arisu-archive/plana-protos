package protos

type MailListSemiPermanentResponse struct {
	ResponsePacket
	MailDBs []MailDB `json:",omitempty,omitzero"`
	Count   int64    `json:",omitempty,omitzero"`
}
