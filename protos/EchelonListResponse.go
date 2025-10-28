package protos

type EchelonListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EchelonDBs []EchelonDB `json:",omitempty,omitzero"`
	ArenaEchelonDB EchelonDB `json:",omitempty,omitzero"`
}
