package protos

type EchelonBatchSaveResponse struct {
	ResponsePacket
	EchelonDBs []*EchelonDB
}
