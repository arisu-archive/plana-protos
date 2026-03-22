package protos

type EchelonListResponse struct {
	ResponsePacket
	EchelonDBs     []EchelonDB
	ArenaEchelonDB EchelonDB
}
