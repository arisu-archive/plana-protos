package protos

type ScenarioWithdrawEchelonResponse struct {
	ResponsePacket
	SaveDataDB         StoryStrategyStageSaveDB
	WithdrawEchelonDBs []EchelonDB
}
