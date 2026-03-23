package protos

type MailListSemiPermanentResponse struct {
	ResponsePacket
	MailDBs []*MailDB
	Count   int64 `json:",omitempty,omitzero"`
}
