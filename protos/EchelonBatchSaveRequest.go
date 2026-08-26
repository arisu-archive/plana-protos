package protos

type EchelonBatchSaveRequest struct {
	RequestPacket
	EchelonDBs []*EchelonDB
}
