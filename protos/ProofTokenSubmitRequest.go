package protos

type ProofTokenSubmitRequest struct {
	RequestPacket
	Answer int64 `json:",omitempty,omitzero"`
}
