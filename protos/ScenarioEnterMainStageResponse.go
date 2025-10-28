package protos

type ScenarioEnterMainStageResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SaveDataDB StoryStrategyStageSaveDB `json:",omitempty,omitzero"`
}
