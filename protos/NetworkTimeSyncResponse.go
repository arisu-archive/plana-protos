package protos

type NetworkTimeSyncResponse struct {
	ResponsePacket
	ReceiveTick int64 `json:",omitempty,omitzero"`
	EchoSendTick int64 `json:",omitempty,omitzero"`
}
