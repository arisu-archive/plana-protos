package protos

type MailListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MailDBs []MailDB `json:",omitempty,omitzero"`
	Count int64 `json:",omitempty,omitzero"`
}
