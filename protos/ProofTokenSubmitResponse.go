package protos

type ProofTokenSubmitResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
