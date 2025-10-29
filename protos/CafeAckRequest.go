package protos

type CafeAckRequest struct {
	RequestPacket
	CafeDBId int64 `json:",omitempty,omitzero"`
}
