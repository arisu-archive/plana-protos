package protos

type ConquestConquerDeployEchelonResponse struct {
	ResponsePacket
	ConquestEchelonDBs []ConquestEchelonDB `json:",omitempty,omitzero"`
	ConquestInfoDB ConquestInfoDB `json:",omitempty,omitzero"`
}
