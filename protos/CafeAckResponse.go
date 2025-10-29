package protos

type CafeAckResponse struct {
	ResponsePacket
	CafeDB CafeDB `json:",omitempty,omitzero"`
}
