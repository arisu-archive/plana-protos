package protos

type ConquestConquerDeployEchelonResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ConquestEchelonDBs []ConquestEchelonDB `json:",omitempty,omitzero"`
	ConquestInfoDB ConquestInfoDB `json:",omitempty,omitzero"`
}
