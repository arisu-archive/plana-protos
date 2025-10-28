package protos

type NetworkTimeSyncResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ReceiveTick int64 `json:",omitempty,omitzero"`
	EchoSendTick int64 `json:",omitempty,omitzero"`
}
