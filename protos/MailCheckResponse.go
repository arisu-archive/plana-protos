package protos

type MailCheckResponse struct {
	ResponsePacket
	Count                  int64 `json:",omitempty,omitzero"`
	CommonMailCount        int64 `json:",omitempty,omitzero"`
	SemiPermanentMailCount int64 `json:",omitempty,omitzero"`
}
