package protos

type EchelonListResponse struct {
	ResponsePacket
	EchelonDBs []EchelonDB `json:",omitempty,omitzero"`
	ArenaEchelonDB EchelonDB `json:",omitempty,omitzero"`
}
