package protos

type MailCheckResponse struct {
	ResponsePacket
	CommonMailCount        int64 `json:",omitempty,omitzero"`
	SemiPermanentMailCount int64 `json:",omitempty,omitzero"`
}
