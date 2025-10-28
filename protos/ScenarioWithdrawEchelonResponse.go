package protos

type ScenarioWithdrawEchelonResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDataDB StoryStrategyStageSaveDB `json:",omitempty,omitzero"`
	WithdrawEchelonDBs []EchelonDB `json:",omitempty,omitzero"`
}
