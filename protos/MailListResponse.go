package protos

type MailListResponse struct {
	ResponsePacket
	MailDBs []MailDB `json:",omitempty,omitzero"`
	Count int64 `json:",omitempty,omitzero"`
}
