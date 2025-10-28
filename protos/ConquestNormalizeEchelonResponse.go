package protos

type ConquestNormalizeEchelonResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ConquestEchelonDB ConquestEchelonDB `json:",omitempty,omitzero"`
}
