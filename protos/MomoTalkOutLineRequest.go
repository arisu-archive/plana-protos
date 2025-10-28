package protos

type MomoTalkOutLineRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
