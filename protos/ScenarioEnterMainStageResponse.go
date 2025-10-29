package protos

type ScenarioEnterMainStageResponse struct {
	ResponsePacket
	SaveDataDB StoryStrategyStageSaveDB `json:",omitempty,omitzero"`
}
