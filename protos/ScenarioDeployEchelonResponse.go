package protos

type ScenarioDeployEchelonResponse struct {
	ResponsePacket
	SaveDataDB StoryStrategyStageSaveDB `json:",omitempty,omitzero"`
}
