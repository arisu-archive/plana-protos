package protos

type ConquestNormalizeEchelonResponse struct {
	ResponsePacket
	ConquestEchelonDB ConquestEchelonDB `json:",omitempty,omitzero"`
}
