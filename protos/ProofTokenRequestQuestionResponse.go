package protos

type ProofTokenRequestQuestionResponse struct {
	ResponsePacket
	Hint int64 `json:",omitempty,omitzero"`
	Question string `json:",omitempty,omitzero"`
	Protocol Protocol `json:",omitempty,omitzero"`
}
