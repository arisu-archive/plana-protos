package protos

type EventContentWithdrawEchelonResponse struct {
	ResponsePacket
	SaveDataDB         EventContentMainStageSaveDB
	WithdrawEchelonDBs []EchelonDB
}
