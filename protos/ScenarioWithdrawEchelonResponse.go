package protos

type ScenarioWithdrawEchelonResponse struct {
	ResponsePacket
	SaveDataDB         *StoryStrategyStageSaveDB `json:",omitempty,omitzero"`
	WithdrawEchelonDBs []*EchelonDB
}
