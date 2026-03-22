package protos

type MailListResponse struct {
	ResponsePacket
	MailDBs []MailDB
	Count   int64 `json:",omitempty,omitzero"`
}
