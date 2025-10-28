package protos

type ProofTokenRequestQuestionRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
