package protos

type MailCheckResponse struct {
	ResponsePacket
	Count int64 `json:",omitempty,omitzero"`
}
