package protos

type MailCheckResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Count int64 `json:",omitempty,omitzero"`
}
