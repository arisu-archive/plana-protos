package protos

type MiniGameResultResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
