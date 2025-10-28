package protos

type EchelonSaveResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EchelonDB EchelonDB `json:",omitempty,omitzero"`
}
