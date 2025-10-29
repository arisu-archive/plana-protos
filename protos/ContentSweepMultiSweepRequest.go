package protos

type ContentSweepMultiSweepRequest struct {
	RequestPacket
	MultiSweepParameters []MultiSweepParameter `json:",omitempty,omitzero"`
}
