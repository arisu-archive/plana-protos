package protos

type CafeAckRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	CafeDBId int64 `json:",omitempty,omitzero"`
}
