package protos

type ConquestConquerDeployEchelonResponse struct {
	ResponsePacket
	ConquestEchelonDBs []*ConquestEchelonDB
	ConquestInfoDB     *ConquestInfoDB `json:",omitempty,omitzero"`
}
