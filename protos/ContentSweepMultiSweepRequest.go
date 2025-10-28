package protos

type ContentSweepMultiSweepRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MultiSweepParameters []MultiSweepParameter `json:",omitempty,omitzero"`
}
